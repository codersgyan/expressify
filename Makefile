BINARY_NAME=expressify
ENTRY_POINT=cmd/expressify/main.go

all: test build

build:
	@go build -o bin/$(BINARY_NAME) $(ENTRY_POINT)

run:
	@go run $(ENTRY_POINT)

test:
	@go test ./...

clean:
	@go clean
	@rm -f bin/$(BINARY_NAME)

.PHONY: all build run test clean
