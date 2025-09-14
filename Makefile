.PHONY: all build clean run

BINARY_NAME = focusmode
BUILD_DIR = build
OUT = $(BUILD_DIR)/$(BINARY_NAME)

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(OUT) .

run: build
	$(OUT)

clean:
	rm -rf $(BUILD_DIR)