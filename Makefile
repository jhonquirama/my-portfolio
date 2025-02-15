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

coverage:
	./scripts/coverage


test-coverage-html:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

test-coverage-total:
	go tool cover -func=coverage.out

.PHONY: lint clean build coverage test-coverage-html test-coverage-total