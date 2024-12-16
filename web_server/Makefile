# Go package name
PACKAGE_NAME := sampleCrud

# Output directory
OUTPUT_DIR := build

# Target architectures
TARGETS := linux/amd64 darwin/amd64 windows/amd64

# Go flags
GOFLAGS := -ldflags="-s -w"

# UPX options
UPXFLAGS := --best --lzma

# Build all targets
all: clean build compress

# Clean build directory
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(OUTPUT_DIR)
	@mkdir -p $(OUTPUT_DIR)

# Build binaries for all targets
build: $(TARGETS)

$(TARGETS):
	@echo "Building for $@..."
	@GOOS=$(word 1,$(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) go build $(GOFLAGS) -o $(OUTPUT_DIR)/$(PACKAGE_NAME)-$(word 1,$(subst /, ,$@))

# Compress binaries with UPX
compress:
	@echo "Compressing binaries with UPX..."
	@for file in $(OUTPUT_DIR)/*; do \
		if [ -f "$$file" ]; then \
			upx $(UPXFLAGS) "$$file"; \
		fi; \
	done

# Default target
.PHONY: all clean build compress
