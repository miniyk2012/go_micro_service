


```bash
protoc --proto_path=proto --go_out=proto \
--go_opt=paths=source_relative \
--go-grpc_out=proto \
--go-grpc_opt=paths=source_relative \
proto/book/book.proto proto/book/price.proto proto/author/author.proto
```
--proto_path有一个别名-I