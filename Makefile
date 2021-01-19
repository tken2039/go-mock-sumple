.PHONY: run
run:
	go run ./cmd/main

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o bin ./...

.PHONY: start
start: build
	./bin/server
