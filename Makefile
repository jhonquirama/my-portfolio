APPNAME ?= my-portfolio

# used by `test` target
export REPORTS_DIR=./reports
# used by lint target
export GOLANGCILINT_VERSION=v1.55.1

test-coverage-html:
	go test ./internal/... -coverprofile=coverage.out && go tool cover -html=coverage.out

test-coverage-total:
	go tool cover -func=coverage.out

start:
	go run cmd/main.go



build: clean swagger
	mkdir -p build
	GOOS=$(GOOS) GOARCH=$(GOARCH) APPNAME=$(APPNAME) ./scripts/build

run: build
	./build/${APPNAME}

test:
	./scripts/unit-test

test-report:
	./scripts/show-tests

lint:
	./scripts/lint

proto-lint:
	./scripts/proto-lint

clean:
	APPNAME=$(APPNAME) ./scripts/clean

swagger:
	mkdir -p doc
	./scripts/swagger

swagger-report:
	./scripts/show-swagger

aws_build:  clean
	mkdir -p build
	GOOS=linux GOARCH=amd64 APPNAME=$(APPNAME) ./scripts/build

aws_zip: aws_build
	zip -j ./build/${APPNAME}.zip ./build/${APPNAME}
	@echo "We have built a ${APPNAME}.zip file, check it"
	ls -lh ./build/${APPNAME}.zip

protoc:
	for i in $(filter-out $@,$(MAKECMDGOALS)); do \
        PROTOC_TARGET=$$i ./scripts/protoc ; \
    done

.PHONY: build run test test-report lint clean zip aws_build protoc