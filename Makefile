.PHONY: run info %
.DEFAULT_GOAL := info

info:
	echo "No target specified: try 'run'"

test:
	go test ./...
