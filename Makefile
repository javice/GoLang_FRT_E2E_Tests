# Makefile

# Variables
BINARY_NAME=GoLang_FRT_E2E_Tests
MAIN_PACKAGE=./cmd/main.go
TEST_REPORT_DIR=reports

# Comandos
.PHONY: all build test clean run lint report

all: clean lint test build

build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

test:
	go test -v ./...

test-report:
	# ============ LIMPIAMOS EL DIRECTORIO DE REPORTES ============
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f $(TEST_REPORT_DIR)/*
	# ============ LIMPIAMOS CACHE ============
	@mkdir -p $(TEST_REPORT_DIR)
	go clean -testcache
	# ============ REALIZAMOS TESTS ============
	go test -v -count=1 ./tests/e2e/... 
	# ============ GENERAMOS REPORTE ============
	go run cmd/generate_report/main.go

clean:
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f $(TEST_REPORT_DIR)/*

run:
	go run $(MAIN_PACKAGE) | tee reports/test.log


lint:
	golangci-lint run

deps:
	go mod download
	go mod tidy