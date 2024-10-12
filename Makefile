# Makefile for Go Project
# Variables
GO=go
PKG=./...
OUTDIR=build/
INSTALL_DIR=/usr/local/bin/
VERSION=1.0.0
LDFLAGS=-s -w

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	@echo "Building the binary..."
	mkdir -p $(OUTDIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(OUTDIR) $(PKG)

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	$(GO) run $(PKG)

# Clean the build
.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -rf $(OUTDIR)

.PHONY: install
install: build
	install -m 755 $(OUTDIR)/lichess-tui $(INSTALL_DIR)

# Show help
.PHONY: help
help:
	@echo "Makefile for Go Project"
	@echo "Available commands:"
	@echo "  make build     - Build the binary"
	@echo "  make install   - Install the binary"
	@echo "  make run       - Run the application"
	@echo "  make clean     - Clean the build"
	@echo "  make help      - Show this help message"
