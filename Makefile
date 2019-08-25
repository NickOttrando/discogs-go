.DEFAULT_GOAL := test

test:
	go test ./discogs/

fmt:
	go fmt ./...

build:
	go build
