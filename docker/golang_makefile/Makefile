GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=app
CURRENT_PLATFORM=darwin

VERSION ?= latest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

PKGS := $(shell go list ./... | grep -v /vendor)

lint:
	golint $(PKGS)

.PHONY: test
test: lint
	$(GOTEST) -v $(PKGS)

.PHONY: run
run: build
	./release/$(BINARY_NAME)-$(VERSION)-$(CURRENT_PLATFORM)-amd64

deps:
	dep ensure
	# $(GOGET) -u github.com/markbates/goth

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 $(GOBUILD) -o release/$(BINARY_NAME)-$(VERSION)-$(os)-amd64

.PHONY: build
build: windows linux darwin

.PHONY: docker
docker: build
	docker build -t go-makefile-test-$(BINARY_NAME) .

clean: 
	$(GOCLEAN)
	rm -rf release