default: out/example
clean:
	rmdir /s /q out
test:
	go vet && go test
define VERSION_BODY
package main

const BuildVersion = "$(shell git describe)"
endef
export VERSION_BODY
out/example: implementation.go main.go version.go
	mkdir out
	@echo "$$VERSION_BODY" > version.go
	go build -o out/example ./
