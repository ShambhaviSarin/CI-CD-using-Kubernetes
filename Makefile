BINARY_NAME=helloworld
IMAGE_NAME=ghcr.io/ShambhaviSarin/helloworld:latest

.PHONY: all build test docker-build clean run

all: build

build:
	go build -o $(BINARY_NAME) ./...

test:
	go test ./...

docker-build:
	docker build -t $(IMAGE_NAME) .

run:
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
