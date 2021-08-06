run:
	go run cmd/main.go

build:
	go build cmd/main.go -o main

lint:
	golint ./...

test:
	go test -cover ./...

test-html: 
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
