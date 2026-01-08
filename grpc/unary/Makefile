# remember to export GOPATH if not set
# export GOPATH=$HOME/go

.PHONY: compile
compile:
	protoc -I./proto \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=. --go_opt=module=unary \
		--go-grpc_out=. --go-grpc_opt=module=unary \
		--grpc-gateway_out=. --grpc-gateway_opt=module=unary \
		./proto/hello.proto
	protoc -I./proto \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=. --go_opt=module=unary \
		--go-grpc_out=. --go-grpc_opt=module=unary \
		./proto/add.proto


.PHONY: alt-compile
alt-compile:
	protoc -I ./proto \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
		-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=plugins=grpc:../ \
		./proto/hello.proto