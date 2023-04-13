package zero_interruption

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/tokopedia/socketmaster/child"
)

var (
	defaultGraceTimeout = 10 * time.Second
)

var (
	// ErrGraceShutdownTimeout happens when the server graceful shutdown exceed the given grace timeout
	ErrGraceShutdownTimeout = errors.New("server shutdown timed out")
)

// ServeHTTP start the http server on the given address and server
func ServeHTTP(port string, srv http.Server) error {
	// add colon
	port = ":" + port

	// start graceful listener
	lis, err := listen(port)
	if err != nil {
		return err
	}

	stoppedCh := waitTermSig(func(ctx context.Context) error {
		stopped := make(chan struct{})

		ctx, cancel := context.WithTimeout(ctx, defaultGraceTimeout)
		defer cancel()

		go func() {
			log.Println("shutting down old server")
			srv.Shutdown(ctx)
			close(stopped)
		}()

		select {
		case <-ctx.Done():
			// server shutdown took longer than timeout, thus the error
			return ErrGraceShutdownTimeout
		case <-stopped:
			// server shutdown finished properly
		}

		return nil
	})

	log.Printf("http server running on address: %v", port)

	// start serving
	if err := srv.Serve(lis); err != http.ErrServerClosed {
		return err
	}

	<-stoppedCh
	log.Println("server stopped")
	return nil
}

// WaitTermSig wait for termination signal and then execute the given handler
// when the signal received
//
// The handler is usually service shutdown, so we can properly shutdown our server upon SIGTERM
//
// It returns channel which will be closed after the signal received and the handler executed.
// We can use the signal to wait for the shutdown to be finished
func waitTermSig(handler func(context.Context) error) <-chan struct{} {
	stoppedCh := make(chan struct{})

	go func() {
		signals := make(chan os.Signal, 1)

		// wait for the sigterm
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		if err := handler(context.Background()); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		} else {
			log.Println("graceful shutdown succeed")
		}

		close(stoppedCh)
	}()

	return stoppedCh
}

// Listen listens to the given port or to file descriptor as specified by socketmaster
//
// This method is taken from tokopedia/grace repo and modified to work with
// socketmaster's -wait-child-notif option
func listen(port string) (net.Listener, error) {
	var l net.Listener

	// see if we run under socketmaster
	fd := os.Getenv("EINHORN_FDS")
	if fd != "" {
		// EINHORN_FDS has value, so we most likely running under socketmaster
		sock, err := strconv.Atoi(fd)
		if err != nil {
			return nil, err
		}

		log.Println("socketmaster detected, listening on", fd)

		file := os.NewFile(uintptr(sock), "listener")
		fl, err := net.FileListener(file)
		if err != nil {
			return nil, err
		}

		l = fl
	}

	if l != nil {
		// we run under socketmaster, thus we notify our service readiness
		notifSocketMaster()
		return l, nil
	}

	// we do not run under socketmaster, thus we normally listen to tcp connections
	return net.Listen("tcp4", port)
}

// notifSocketMaster notify socket master about our readyness
func notifSocketMaster() {
	go func() {
		err := child.NotifyMaster()
		if err != nil {
			log.Printf("failed to notify socketmaster: %v, ignore if you don't use `wait-child-notif` option", err)
		} else {
			log.Println("successfully notify socketmaster")
		}
	}()
}
