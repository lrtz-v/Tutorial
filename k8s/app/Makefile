BINARY_NAME=app

VERSION ?= "1.0.0"
PKGS := $(shell go list ./...)

lint:
	golangci-lint run ./...

test:
	go test -v $(PKGS)

run:
	go build -o $(BINARY_NAME) . && ./$(BINARY_NAME)

image:
	docker build -t $(BINARY_NAME):$(VERSION) .

clean:
	rm app
