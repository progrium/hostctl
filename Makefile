NAME=hostctl
ARCH=$(shell uname -m)
VERSION=0.2.0dev


.PHONY: build
build: local/bin/hostctl

local/bin/hostctl: *.go **/*.go
	go build -o local/bin/hostctl .

.PHONY: test
test:
	go test -v
