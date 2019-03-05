
### Compile

```
protoc -I/home/kfoozminus/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        -I./src --go_out=plugins=grpc:./build/gen \
        --grpc-gateway_out=logtostderr=true:./build/gen \
        src/foo.proto
```


### Run gRPC server

```
go run cmd/server/main.go
```

### Run gRPC client

```
go run cmd/client/main.go
```

### HTTP/1 Req

```
curl "http://localhost:4040/hello?name=Jenny"
```


