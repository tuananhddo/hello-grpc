.PHONY: build install-tools

GRPC_GATEWAY := $(shell go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2)
PROTOC_INCLUDES := .:${GRPC_GATEWAY}/third_party/googleapis

build:
	go build -o server cmd/server/server.go
	go build -o client cmd/client/client.go
	go build -o client cmd/gateway/gateway.go

install-tools:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

gen:
	cd proto && \
		protoc -I${PROTOC_INCLUDES} \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--grpc-gateway_out=paths=source_relative,logtostderr=true:. hello.proto
	cd proto && \
		protoc -I${PROTOC_INCLUDES} \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--grpc-gateway_out=paths=source_relative,logtostderr=true:. example.proto