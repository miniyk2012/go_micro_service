# lesson28

两个关键点：
1. gRPC客户端如何拿到所有的服务机地址列表
2. gRPC拿到服务机地址列表后如何决定连哪个？


```bash
protoc -I=pb \
   --go_out=pb --go_opt=paths=source_relative \
   --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
   hello.proto
```