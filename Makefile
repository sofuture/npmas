PROGRAM = npmas
PREFIX  = bin

default: build

build: generate
	@godep go build -a -ldflags "-X github.com/sofuture/npmas/version.GitCommit=$$(git rev-parse HEAD)" -o $(PREFIX)/$(PROGRAM)
	@cp $(PREFIX)/$(PROGRAM) $(GOPATH)/$(PREFIX)/$(PROGRAM) || true

test: generate
	@godep go test ./... -parallel=4 -race

generate:
	@godep go generate ./...

clean:
	@rm -rf bin/*

.PHONY: default build generate test clean
