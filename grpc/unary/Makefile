.PHONY: compile
compile:
	protoc -I./proto \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=. --go_opt=module=unary \
		--go-grpc_out=. --go-grpc_opt=module=unary \
		./proto/*.proto