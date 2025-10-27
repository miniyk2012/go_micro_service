.PHONY: gen server client


PROTO_DIR=routeguide

gen:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(PROTO_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_DIR) \
	--go-grpc_opt=paths=source_relative \
	$(PROTO_DIR)/route_guide.proto


server:
	go run server/server.go

client:
	go run client/client.go