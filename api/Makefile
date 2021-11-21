.PHONY: setup install mockgen protoc fmt vet lint test build

GOLANGCI_VERSION := 1.43.0

LINT_PACKAGES := $(shell go list $(CURDIR)/... | grep -v -e "mock" -v -e "proto" -v -e "third_party" -v -e "tmp")
TEST_PACKAGES := $(shell go list $(CURDIR)/internal/... $(CURDIR)/pkg/...)
EXTERNAL_PACKAGE=$(CURDIR)/third_party

setup: install
	wget -O - -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}

install: install-third-party
	go install github.com/bazelbuild/bazelisk@v1.10.1
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
	go install github.com/golang/mock/mockgen@v1.6.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

install-third-party:
	rm -rf ${EXTERNAL_PACKAGE}/**/*
	mkdir -p ${EXTERNAL_PACKAGE}/github.com/envoyproxy
	git clone https://github.com/envoyproxy/protoc-gen-validate.git ${EXTERNAL_PACKAGE}/github.com/envoyproxy/protoc-gen-validate

mockgen:
	rm -rf ./mock
	go generate ./...
	./hack/build-mocks.sh

protoc:
	find ./proto/**/*.go -type f | xargs rm -rf
	protoc \
		-I ./proto \
		-I ./third_party \
		--go_out ./proto \
		--go_opt paths=source_relative \
		--go-grpc_out ./proto \
		--go-grpc_opt paths=source_relative \
		--validate_out lang=go:./proto \
		--validate_opt paths=source_relative \
		./proto/**/*.proto

fmt:
	! gofmt -d -s ./cmd ./config ./hack ./internal ./pkg | grep '^'

vet:
	go vet $(LINT_PACKAGES)

lint:
	./bin/golangci-lint run -c .golangci.yaml ./...

test:
	go test -v -cover -coverprofile=coverage.txt -covermode=atomic $(TEST_PACKAGES)

build:
	go build -o ./app ./cmd/${SERVICE}/main.go

dev:
	$(MAKE) build SERVICE=${SERVICE}
	./app
