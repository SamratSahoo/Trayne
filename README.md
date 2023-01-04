# Trayne

## Getting Started

Start the peripheral node

```
go run main.go --type "peripheral"
```

Start the orchestrator node & specify IP and port of the peripheral node(s)

```
go run main.go --type "orchestrator" --peers "127.0.0.1:4000"
```

Send a message to the orchestrator to initialize training

```
netcat 127.0.0.1 3000
{"messageType":"orchestrator training init"}
```

Alternatively, run the project as a client:

```
go run main.go --type client --host 127.0.0.1 --port 3000
```
