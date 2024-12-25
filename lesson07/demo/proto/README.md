


```bash
protoc --proto_path=proto --go_out=proto \
--go_opt=paths=source_relative book/price.proto
```
--proto_path有一个别名-I