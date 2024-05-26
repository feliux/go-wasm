all: help

BIN_NAME = json.wasm
BIN_FOLDER := assets
COVERAGE_FILE = coverage.out
REPORT_FOLDER = reports

GO_LDFLAGS = -ldflags "-s -w"

help:
	@echo '	Usage:'
	@echo '	  author	Display author.'
	@echo '	  setup		Setup development environment.'
	@echo '	  lint		Download Go packages, format and vet.'
	@echo '	  test		Execute tests.'
	@echo '	  install-lint	Download and install golangci-lint.'
	@echo '	  static-check	Execute golangci-lint.'
	@echo '	  build		Compile to binary.'
	@echo '	  copy-hooks	Configure git hooks.'
	@echo '	  clean		Clean binay file.'

author:
	@echo "	Project by feliux"
	@echo "	https://github.com/feliux"

setup: install-lint #copy-hooks
	cp $(GOROOT)/misc/wasm/wasm_exec.js $(BIN_FOLDER)

lint:
	go get ./...
	go fmt ./...
	go vet ./...

test: lint
	go test ./... -coverprofile=$(REPORT_FOLDER)/$(COVERAGE_FILE)
	go tool cover -func $(REPORT_FOLDER)/$(COVERAGE_FILE) | grep "total:" | awk '{ print ((int($$3) > 80) != 1) }'
	go tool cover -html=$(REPORT_FOLDER)/$(COVERAGE_FILE) -o $(REPORT_FOLDER)/coverage.html

install-lint:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.58.2

static-check:
	golangci-lint run --issues-exit-code 0

build: lint test static-check
	GOOS=js GOARCH=wasm go build -o $(BIN_FOLDER)/${BIN_NAME} ./cmd/wasm/main.go

copy-hooks:
	chmod +x scripts/hooks/pre-commit
	cp -r scripts/hooks/* .git/hooks/

clean:
	rm -f $(BIN_FOLDER)*
