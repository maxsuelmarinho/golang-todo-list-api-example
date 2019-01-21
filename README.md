# Golang Todo list API example

**Dependencies:**
```
> go mod init github.com/maxsuelmarinho/golang-todo-list-api-example

> go get -u github.com/golang/protobuf/protoc-gen-go

> go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

> go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

```
> cp -rv <protobuffer-installation-dir>/include/* third_party/

# gRPC Gateway dependencies:
> cp -rv $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google third_party/google
> cp $GOPATH/src/github.com/grpc-ecosystem/protoc-gen-swagger/options/annotations.proto third_party/protoc-gen-swagger/options/
> cp $GOPATH/src/github.com/grpc-ecosystem/protoc-gen-swagger/options/openapiv2.proto third_party/protoc-gen-swagger/options/
```

**Start Server:**
```
> cd cmd/server
> go build
> server -grpc-port=9090 -http-port=8080 -db-host=<host>:3306 -db-user=<user> -db-password=<password> -db-schema=<schema>
```

**Start Client:**
```
# gRPC client
> cd cmd/client-grpc
> go build
> client-grpc -server=localhost:9090

# REST client
> cd cmd/client-rest
> go build
> client-rest -server=http://localhost:8080
```