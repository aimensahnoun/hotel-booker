build:
	go build -o bin/api

dev:
	go run main.go

run:
	./bin/api

test:
	go test -v ./...