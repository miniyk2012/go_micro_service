.PHONY: gen server client


PROTO_DIR=proto

gen:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(PROTO_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_DIR) \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative \
	$(PROTO_DIR)/hello.proto


server:
	go run server/main.go

client:
	go run client/main.go -name abc