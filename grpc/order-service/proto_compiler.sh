protoc -I proto/ proto/*.proto --go_out=plugins=grpc:server/ecommerce/
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:client/ecommerce/

