BASE_PATH       = ./
MAIN_GO_PATH    = /main.go

default: run

build:
	go build main.go

install:
	go mod download

test:
	go test ./...

coverage:
	go test -cover ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

just-run:
	go run app.go run-system

run: install just-run