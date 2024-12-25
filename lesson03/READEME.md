# 第一个gRPC示例

 hello world

## 三个步骤
1.编写protobuf文件
2.生成代码(服务端和客户端)
```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
pb/*.proto
```
3.编写业务逻辑代码


## 作业
自己动手写一个 grpc版本的  add(x, y int) int


## python client
```python
## --pyi_out=.是生成stub来自动提示的, 这是因为python的pb是运行时动态获取, 不能直接引用, 因此需要stub
python -m grpc_tools.protoc -Ipb --python_out=. --pyi_out=. --grpc_python_out=. pb/add.proto 
```