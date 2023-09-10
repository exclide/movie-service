.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: run
run:
	go run ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run

.PHONY: protogen
protogen:
	protoc --go_out=. --go-grpc_out=. api/proto/stringer.proto