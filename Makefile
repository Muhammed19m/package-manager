run:
	go run cmd/pm/main.go

build:
	go build -o bin/pm cmd/pm/main.go

test:
	go test -v ./...