BINARY_NAME=app

version = $(v)
ifndef version
	version = "1.0.0"
endif

PKGS := $(shell go list ./...)

lint:
	golangci-lint run ./...

test:
	go test -v $(PKGS)

run:
	go build -o $(BINARY_NAME) . && ./$(BINARY_NAME)

image:
	docker build -t $(BINARY_NAME):$(version) .

clean:
	rm app
