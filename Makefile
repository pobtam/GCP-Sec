BINARY    := gcp-security-analyzer
MAIN      := ./main.go
BUILD_DIR := ./dist
VERSION   := 1.0.0
LDFLAGS   := -ldflags="-X main.version=$(VERSION) -s -w"

.PHONY: all build test test-cover lint clean install run-sample help

all: build

## build: Compile the binary
build:
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) $(MAIN)
	@echo "✓ Built $(BUILD_DIR)/$(BINARY)"

## install: Install the binary to GOPATH/bin
install:
	go install $(LDFLAGS) .
	@echo "✓ Installed $(BINARY)"

## test: Run all unit tests
test:
	go test ./... -v -count=1

## test-cover: Run tests with coverage report
test-cover:
	go test ./... -coverprofile=coverage.out -covermode=atomic
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out | grep total
	@echo "✓ Coverage report written to coverage.html"

## bench: Run benchmark tests
bench:
	go test ./... -bench=. -benchmem -run='^$$'

## lint: Run go vet and staticcheck
lint:
	go vet ./...
	@which staticcheck > /dev/null 2>&1 && staticcheck ./... || echo "staticcheck not installed, skipping"

## fmt: Format all Go source files
fmt:
	gofmt -w .

## tidy: Tidy go.mod and go.sum
tidy:
	go mod tidy

## clean: Remove build artifacts
clean:
	rm -rf $(BUILD_DIR) coverage.out coverage.html

## run-sample: Analyze the sample CSV and generate a report
run-sample: build
	$(BUILD_DIR)/$(BINARY) analyze testdata/sample-findings.csv \
		--output reports/sample-report.md \
		--include-remediation \
		--include-compliance \
		--verbose
	@echo "✓ Report written to reports/sample-report.md"

## run-all-formats: Generate reports in all formats from the sample CSV
run-all-formats: build
	@mkdir -p reports
	$(BUILD_DIR)/$(BINARY) analyze testdata/sample-findings.csv \
		--output-dir reports \
		--formats markdown,json,html,csv \
		--include-remediation \
		--include-compliance

## stats-sample: Show statistics for the sample CSV
stats-sample: build
	$(BUILD_DIR)/$(BINARY) stats testdata/sample-findings.csv

## help: Show this help
help:
	@echo "Available targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'
