.PHONY: setup test install

default: test install

setup:
	go get -d -u github.com/cortesi/modd/cmd/modd
	go get -d -u honnef.co/go/tools/cmd/staticcheck

test:
	go mod tidy
	staticcheck ./...
	go test -timeout=60s ./...

install:
	go install ./cmd/...