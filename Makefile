build:
	@go build -o bin/api

dev:
	@go run main.go

run: build
	@./bin/api

test:
	@go test -v ./...