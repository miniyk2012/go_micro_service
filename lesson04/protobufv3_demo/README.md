# 只生成pb, 不生成rpc

```bash
protoc -I=. --go_out=. --go_opt=paths=source_relative pb/demo.proto
```