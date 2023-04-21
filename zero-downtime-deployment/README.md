# Zero Downtime Deployment

Socketmaster is an application where we can put our apps as its process child. Basically it helps us reload the server without losing current and incoming connections.

We will try to do simple service deployment while keeping our service alive (all the time) under Socketmaster.


### Part 1: Run the Server

Open your terminal, and run the following command:
```
~ ./bin/server
```

Server will prepare itself and notify us if its ready:
```
2023/04/13 15:00:00 initiate new server
2023/04/13 15:00:00 connecting to database...
2023/04/13 15:00:01 database connected
2023/04/13 15:00:01 building cache...
2023/04/13 15:00:02 cache built
2023/04/13 15:00:02 ping upstream service...
2023/04/13 15:00:03 upstream ready
2023/04/13 15:00:03 initiate server finished
2023/04/13 15:00:03 http server running on address: :8000
```

You can check if the server running properly by executing curl command against `localhost:8000`:
```
~ curl localhost:8000
Make it Happen
```


### Part 2: Run the Server under Socketmaster

Terminate previous server by using `Ctrl-c`. Our server already implemented graceful shutdown, so on termination, you will see following messages:
```
2023/04/13 15:00:10 shutting down old server
2023/04/13 15:00:10 graceful shutdown succeed
2023/04/13 15:00:10 server stopped
```

Once done, we can re-run the server under Socketmaster using following command:
```
~ ./bin/socketmaster -command=./bin/server -listen tcp://:8000 -wait-child-notif=true
```

Running our server under Socketmaster require us to do some extra steps:
```
socketmaster[73840] 2023/04/13 15:10:08 Listening on tcp://:8000 <-- extras
socketmaster[73840] 2023/04/13 15:10:08 Starting ./bin/server [./bin/server] <-- extras
socketmaster[73840] 2023/04/13 15:10:08 [73841] 2023/04/13 15:10:08 initiate new server
socketmaster[73840] 2023/04/13 15:10:08 [73841] 2023/04/13 15:10:08 connecting to database...
socketmaster[73840] 2023/04/13 15:10:09 [73841] 2023/04/13 15:10:09 database connected
socketmaster[73840] 2023/04/13 15:10:09 [73841] 2023/04/13 15:10:09 building cache...
socketmaster[73840] 2023/04/13 15:10:10 [73841] 2023/04/13 15:10:10 cache built
socketmaster[73840] 2023/04/13 15:10:10 [73841] 2023/04/13 15:10:10 ping upstream service...
socketmaster[73840] 2023/04/13 15:10:11 [73841] 2023/04/13 15:10:11 upstream ready
socketmaster[73840] 2023/04/13 15:10:11 [73841] 2023/04/13 15:10:11 initiate server finished
socketmaster[73840] 2023/04/13 15:10:11 [73841] 2023/04/13 15:10:11 socketmaster detected, listening on 3 <-- extras
socketmaster[73840] 2023/04/13 15:10:11 [73841] 2023/04/13 15:10:11 http server running on address: :8000
socketmaster[73840] 2023/04/13 15:10:11 [73841] 2023/04/13 15:10:11 successfully notify socketmaster <-- extras
socketmaster[73840] 2023/04/13 15:10:14 Failed to kill old process, because there's no one left in the group
```

These extras were covered in sample code (you can check them on `/zero_interruption/zero_interruption.go`). Also, notice on our example log above that:
- Socketmaster were running with PID 73840, and
- server were running with PID 73841

You can check if the server running properly by executing curl command against `localhost:8000`:
```
~ curl localhost:8000
Make it Happen
```


### Part 3: Make Change on Existing Server

Open `/bin/server.go`. Change the returned string from `Make it Happen` to `Make it Better`:
```
33        // io.WriteString(res, "Make it Happen\n")
34        io.WriteString(res, "Make it Better\n")
```

### Part 4: Deploy New Server

Pretend you are Jenkins. An engineer made some change (Part 3) and now wants to deploy them.

We can start by building the code and replace the old binary:
```
go build -o ./bin/server ./server/server.go
```

Previously, we run the server under Socketmaster (Part 2) with defined command `./bin/server`. We will ask Socketmaster to:
1. spawn new process with same command,
2. wait for new process to be ready, and
3. once new process ready, kill the old one

All can be done by sending SIGHUP to Socketmaster (https://github.com/tokopedia/socketmaster#how-it-works).

To do so, open a new terminal tab. prepare your Socketmaster PID, and run this in your terminal:
```
~ kill -s SIGHUP <YOUR-SOCKETMASTER-PID>
```

Go back to your Socketmaster tab. You will notice some new logs appear:
```
socketmaster[73840] 2023/04/13 15:22:37 Starting ./bin/server [./bin/server]
socketmaster[73840] 2023/04/13 15:22:37 [75255] 2023/04/13 15:22:37 initiate new server
socketmaster[73840] 2023/04/13 15:22:37 [75255] 2023/04/13 15:22:37 connecting to database...
socketmaster[73840] 2023/04/13 15:22:38 [75255] 2023/04/13 15:22:38 database connected
socketmaster[73840] 2023/04/13 15:22:38 [75255] 2023/04/13 15:22:38 building cache...
socketmaster[73840] 2023/04/13 15:22:39 [75255] 2023/04/13 15:22:39 cache built
socketmaster[73840] 2023/04/13 15:22:39 [75255] 2023/04/13 15:22:39 ping upstream service...
socketmaster[73840] 2023/04/13 15:22:40 [75255] 2023/04/13 15:22:40 upstream ready
socketmaster[73840] 2023/04/13 15:22:40 [75255] 2023/04/13 15:22:40 initiate server finished
socketmaster[73840] 2023/04/13 15:22:40 [75255] 2023/04/13 15:22:40 socketmaster detected, listening on 3
socketmaster[73840] 2023/04/13 15:22:40 [75255] 2023/04/13 15:22:40 http server running on address: :8000
socketmaster[73840] 2023/04/13 15:22:40 [75255] 2023/04/13 15:22:40 successfully notify socketmaster
socketmaster[73840] 2023/04/13 15:22:43 [73841] 2023/04/13 15:22:43 shutting down old server
socketmaster[73840] 2023/04/13 15:22:43 [73841] 2023/04/13 15:22:43 graceful shutdown succeed
socketmaster[73840] 2023/04/13 15:22:43 [73841] 2023/04/13 15:22:43 server stopped
socketmaster[73840] 2023/04/13 15:22:43 73841 exit status 0 <nil>
```

You can see that Socketmaster spawning new server with PID 75255. After new server send ready signal to Socketmaster, it then terminate the old server (PID 73841).

Make sure that the new version of binary is running by executing curl command against `localhost:8000`:
```
~ curl localhost:8000
Make it Better
```

You've finished your deployment without downtime!
