.PHONY: compile
compile:
	protoc -I./proto --go_out=. --go_opt=module=unary --go-grpc_out=. --go-grpc_opt=module=unary ./proto/dummy.proto