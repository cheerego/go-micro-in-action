

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
go get -u github.com/golang/protobuf/protoc-gen-go 


protoc  -I. -I /Users/hkn/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis  --go_out=plugins=grpc:. --micro_out=.  greeter.proto

protoc -I/usr/local/include/google/googleapis -I. --go_out=. --micro_out=.  greeter.proto



```