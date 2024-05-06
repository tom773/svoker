# Define variables for directories and target
TARGET := wasm32-unknown-unknown
RELEASE_TARGET := target/$(TARGET)/release
WASM_FILE := $(RELEASE_TARGET)/pokerutil.wasm
OUT_DIR := ../../src/lib/pkg

# Default target
all: bindgen

# Build the wasm file using cargo
build: 
	@echo "Building wasm module..."
	cargo build --release --target=$(TARGET)

# Generate bindings using wasm-bindgen
bindgen: $(WASM_FILE)
	@echo "Generating wasm bindings..."
	wasm-bindgen $(WASM_FILE) --out-dir $(OUT_DIR)

# File dependency
$(WASM_FILE): build

# Clean target to remove generated files
clean:
	@echo "Cleaning up..."
	rm -rf $(RELEASE_TARGET)
	rm -rf $(OUT_DIR)

.PHONY: all build bindgen clean
