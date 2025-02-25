# Makefile

# Variables
BINARY_NAME=GoLang_FRT_E2E_Tests
MAIN_PACKAGE=./cmd/main.go
TEST_REPORT_DIR=reports
GO=go 
GOTEST=$(GO) test
GOMOD=$(GO) mod
GOGET=$(GO) get

# Colores para la salida en consola
CYAN=\033[0;36m
RESET=\033[0m


# Reglas
.DEFAULT_GOAL := all


# Comandos
.PHONY: all build test clean run lint report install

all: clean lint test build

install:
	go mod tidy

build:
	@echo "$(CYAN)Compilando $(BINARY_NAME)$(RESET)"
	go build -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

test:
	@echo "$(CYAN)Limpiando cache de tests$(RESET)"
	go clean -testcache
	@echo "$(CYAN)Ejecutando tests$(RESET)"
	# go test -v ./...
	# go test -v -count=1 ./tests/e2e/... 
	PLAYWRIGHT_BROWSERS_PATH=0 xvfb-run --auto-servernum --server-args='-screen 0 1920x1080x24' go test -v -count=1 ./tests/e2e/... -args --no-sandbox

test-report:
	# ============ LIMPIAMOS EL DIRECTORIO DE REPORTES ============
	@echo "$(CYAN)Limpiando el directorio de reportes$(RESET)"
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f $(TEST_REPORT_DIR)/*
	@mkdir -p $(TEST_REPORT_DIR)
	# ============ LIMPIAMOS CACHE ============
	@echo "$(CYAN)Limpiando cache de tests$(RESET)"
	go clean -testcache
	# ============ REALIZAMOS TESTS ============
	@echo "$(CYAN)Ejecutando tests$(RESET)"
	# go test -v -count=1 ./tests/e2e/... 
	PLAYWRIGHT_BROWSERS_PATH=0 xvfb-run --auto-servernum --server-args='-screen 0 1920x1080x24' go test -v -count=1 ./tests/e2e/... -args --no-sandbox
	# ============ GENERAMOS REPORTE ============
	@echo "$(CYAN)Generando reporte$(RESET)"
	go run cmd/generate_report/main.go

clean:
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f $(TEST_REPORT_DIR)/*

run:
	@echo "$(CYAN)Ejecutando $(BINARY_NAME)$(RESET)"
	go run $(MAIN_PACKAGE) | tee reports/test.log


lint:
	golangci-lint run

deps:
	@echo "$(CYAN)Instalando dependencias$(RESET)"
	go mod download
	go mod tidy