build:
	@go build -o bin/api

dev:
	@go run main.go

run: build
	@./bin/api

seed: 
	@go run ./scripts/seed.go

test:
	@go test -v ./...