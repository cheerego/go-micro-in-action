

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
go get -u github.com/golang/protobuf/protoc-gen-go 

// 直接在grpc-gateway中引入
protoc  -I. -I /Users/hkn/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis --grpc-gateway_out=. --go_out=plugins=grpc:. --micro_out=.  greeter.proto

// ln -s /Users/hkn/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis /usr/local/include/google/googleapis
protoc -I/usr/local/include/google/googleapis  -I.  --grpc-gateway_out=. --go_out=. --micro_out=.  greeter.proto
我们可以把grpc-gateway包下的

把 /Users/hkn/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis/google 拷贝到当前目录
protoc s  -I.  --grpc-gateway_out=. --go_out=. --micro_out=.  greeter.proto


```