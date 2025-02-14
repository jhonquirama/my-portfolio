APPNAME ?= my-portfolio

# used by `test` target
export REPORTS_DIR=./reports
# used by lint target
export GOLANGCILINT_VERSION=v1.55.1

lint:
	./scripts/lint

clean:
	APPNAME=$(APPNAME) ./scripts/clean

build: clean
	mkdir -p build
	GOOS=$(GOOS) GOARCH=$(GOARCH) APPNAME=$(APPNAME) ./scripts/build

.PHONY: lint build