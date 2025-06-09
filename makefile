.PHONY: run test lint tidy

run:          ## go run the service
	go run ./cmd/server

test:         ## unit tests
	go test ./...

lint:         ## static analysis
	golangci-lint run

tidy:         ## clean up go.mod/go.sum
	go mod tidy
