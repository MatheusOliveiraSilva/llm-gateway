.PHONY: hello test-go vet tidy

hello:
	go run ./cmd/hello $(ARGS)

test-go:
	go test ./... -cover

vet:
	go vet ./...

tidy:
	go mod tidy
