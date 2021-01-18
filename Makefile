.PHONY: run
run:
	go run ./cmd/main

.PHONY: test
test:
	go test ./...
