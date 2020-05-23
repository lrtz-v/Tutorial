
# Commands

## protoc compiler

```shell
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:server/ecommerce/
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:client/ecommerce/
```

## golang code compiler

```shell
go build
```

## start service

```shell
./server
./client
```
