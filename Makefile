NAME=xiny
VERSION=$(shell cat VERSION)
BUILD=$(VERSION)-$(shell git rev-parse --short HEAD)
LD_FLAGS="-w -X main.version=$(BUILD)"

clean:
	rm -rf _build/ release/

build:
	dep ensure
	CGO_ENABLED=0 go build -tags release -ldflags $(LD_FLAGS) -o $(NAME)

build-all:
	mkdir -p _build
	GOOS=darwin GOARCH=amd64 go build -tags release -ldflags $(LD_FLAGS) -o _build/$(NAME)-$(VERSION)-darwin-amd64
	GOOS=linux  GOARCH=amd64 go build -tags release -ldflags $(LD_FLAGS) -o _build/$(NAME)-$(VERSION)-linux-amd64
	cd _build; sha256sum * > sha256sums.txt

image:
	docker build -t xiny -f Dockerfile .

release:
	mkdir release
	go get github.com/progrium/gh-release/...
	cp _build/* release
	cd release; sha256sum --quiet --check sha256sums.txt
	gh-release create bcicen/$(NAME) $(VERSION) \
		$(shell git rev-parse --abbrev-ref HEAD) $(VERSION)

.PHONY: build
