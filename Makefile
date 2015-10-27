build:
	go get ./...

all: build test vet

test:
	go test ./...

vet:
	go vet ./...
