
# Commands

## protoc compiler

```shell
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:server/ecommerce/
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:client/ecommerce/
```

## golang code compiler

```shell
go build -i -v -o bin/server
go build -i -v -o bin/client
```

## start service

```shell
./bin/server
./bin/client
```
