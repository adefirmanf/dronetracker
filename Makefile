

.PHONY: clean all init generate generate_mocks

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<

clean:
	rm -rf generated

init: clean generate
	go mod tidy
	go mod vendor

test:
	go clean -testcache
	go test -short -coverprofile coverage.out -short -v ./...

test_api:
	go clean -testcache
	go test ./tests/...

generate: generated generate_mocks

generated: api.yml
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --package generated -generate types,server,spec $< > generated/api.gen.go

# Find interfaces.go files in repository and service folders
INTERFACES_GO_FILES_REPO := $(shell find repository -name "interfaces.go")
INTERFACES_GO_FILES_SERVICE := $(shell find service -name "interfaces.go")

# Generate mock filenames for repository and service interfaces
INTERFACES_GEN_GO_FILES_REPO := $(INTERFACES_GO_FILES_REPO:%.go=%.mock.gen.go)
INTERFACES_GEN_GO_FILES_SERVICE := $(INTERFACES_GO_FILES_SERVICE:%.go=%.mock.gen.go)

# Combine both mock filenames into one variable
INTERFACES_GEN_GO_FILES := $(INTERFACES_GEN_GO_FILES_REPO) $(INTERFACES_GEN_GO_FILES_SERVICE)

# Target to generate mocks
generate_mocks: $(INTERFACES_GEN_GO_FILES)

# Rule to generate mock for each interface
%.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))