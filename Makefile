NAME=xiny
VERSION=$(shell cat VERSION)
BUILD=$(shell git rev-parse --short HEAD)
LD_FLAGS="-w -X main.version=$(VERSION) -X main.build=$(BUILD)"

clean:
	rm -rf _build/ release/

build:
	go mod download
	CGO_ENABLED=0 go build -tags release -ldflags $(LD_FLAGS) -o $(NAME) ./cmd/xiny

build-dev:
	CGO_ENABLED=0 go build -ldflags "-w -X main.version=dev-build -X main.build=$(BUILD)" -o $(NAME) ./cmd/xiny

build-all:
	mkdir -p _build
	GOOS=darwin GOARCH=amd64 go build -tags release -ldflags $(LD_FLAGS) -o _build/$(NAME)-$(VERSION)-darwin-amd64 ./cmd/xiny
	GOOS=linux  GOARCH=amd64 go build -tags release -ldflags $(LD_FLAGS) -o _build/$(NAME)-$(VERSION)-linux-amd64 ./cmd/xiny
	cd _build; sha256sum * > sha256sums.txt

image:
	docker build -t xiny -f Dockerfile .

release:
	mkdir release
	cp _build/* release
	cd release; sha256sum --quiet --check sha256sums.txt && \
	gh release create $(VERSION) -d -t v$(VERSION) *

.PHONY: build
