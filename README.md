## Implementation

The code follows the specification that was provided with one design decision on the part of the client.
The client is responsible for providing it's username as metadata in the Context in subsequent calls after it's first call to Connect inorder to map the client to a username.

## Prerequisites
### Generate protogen
```bash
protoc -I=./proto --go_out=. ./proto/contract.proto --go-grpc_out=. 
```

### Run server with it's associated client example
```bash
(go run server/server.go) & (go run client/client.go)
```

The logs in `logs.txt` can be reproduced from the above command.