.PHONY: gen server client start stop status restart

PROTO_DIR=pb

gen:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(PROTO_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_DIR) \
	--go-grpc_opt=paths=source_relative \
	$(PROTO_DIR)/hello.proto

server:
	go run server/main.go -port=8972
	go run server/main.go -port=8973
	go run server/main.go -port=8974

client:
	go run client/*.go

# 启动所有服务器
start:
	nohup go run server/main.go -port=8972 > server1.log 2>&1 & echo $$! > server1.pid
	nohup go run server/main.go -port=8973 > server2.log 2>&1 & echo $$! > server2.pid
	nohup go run server/main.go -port=8974 > server3.log 2>&1 & echo $$! > server3.pid
	@echo "Servers started. PIDs saved to server{1,2,3}.pid"

# 停止所有服务器
stop:
	@kill `cat server1.pid` 2>/dev/null || true
	@kill `cat server2.pid` 2>/dev/null || true
	@kill `cat server3.pid` 2>/dev/null || true
	@rm -f server*.pid
	@echo "Servers stopped."

# 查看服务器状态
status:
	@echo "Server 1 (port 8972): $$([ -f server1.pid ] && kill -0 `cat server1.pid` 2>/dev/null && echo "Running" || echo "Stopped")"
	@echo "Server 2 (port 8973): $$([ -f server2.pid ] && kill -0 `cat server2.pid` 2>/dev/null && echo "Running" || echo "Stopped")"
	@echo "Server 3 (port 8974): $$([ -f server3.pid ] && kill -0 `cat server3.pid` 2>/dev/null && echo "Running" || echo "Stopped")"

# 重启所有服务器
restart: stop start