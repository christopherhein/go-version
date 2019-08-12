.PHONY: all
all: build-example test

.PHONY: test
test:
	go test -v ./...

.PHONY: build-example
build-example:
	./hack/build_example.sh