.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: run
run:
	go run ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: protogen
protogen:
	protoc --go_out=api/proto --go-grpc_out=api/proto api/proto/stringer.proto

.PHONY: migr
migr:
	migrate -path migrations/ -database "$MIGRATE_DB" up