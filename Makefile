generate:
	go generate ./...
.PHONY: generate

examples:
	cd examples && go generate
.PHONY: examples

tidy:
	go mod tidy
.PHONY: tidy

tidy_examples:
	cd examples && go mod tidy
.PHONY: tidy_examples

tidy_all: tidy tidy_examples
