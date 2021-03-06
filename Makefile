.PHONY: setup test install

default: test install

setup:
	go get -u github.com/cortesi/modd/cmd/modd
	go get -u honnef.co/go/tools/cmd/staticcheck

test:
	staticcheck ./...
	go test -timeout=60s ./...

install:
	go install ./cmd/...