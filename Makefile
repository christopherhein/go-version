.PHONY: all
all: build-example test

.PHONY: lint
lint:
	@golint -set_exit_status ./...

test-fmt:
	./hack/test-fmt.sh

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: build-example
build-example:
	./hack/build_example.sh