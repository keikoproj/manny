BINARY := manny
PKGS := $(shell go list ./... | grep -v /vendor)

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all build clean spotless test cover release artifactory-upload ${GOBIN}/${BINARY}

all: vendor test build

build: ${GOBIN}/${BINARY}

${GOBIN}/${BINARY}:
	go build -o $@ main.go

clean:
	@echo "Removing package object files..."
	@go clean ${PKGS}
	@echo "Removing cache test results..."
	@go clean -testcache

spotless: clean
	@echo "Removing vendor directory..."
	@-rm -rf vendor

vendor: spotless
	@echo "Refreshing dependencies..."
	@go mod tidy && go mod vendor

test:
	go test ${PKGS} ${TESTARGS}

cover: TESTARGS=-coverprofile=coverage.out
cover: test
	go tool cover -func=coverage.out -o coverage.txt
	go tool cover -html=coverage.out -o coverage.html
	@cat coverage.txt
	@echo "Run 'open coverage.html' to view coverage report."

VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

release: windows linux darwin
