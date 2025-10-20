```bash
## --pyi_out=.是生成stub来自动提示的, 这是因为python的pb是动态代码, 没有类可以直接引用, 因此需要stub
uv run python -m grpc_tools.protoc -Ipb --pyi_out=. --python_out=. --grpc_python_out=. pb/*.proto
```