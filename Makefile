BINARY := wt
DIST_DIR := dist
OUT := $(DIST_DIR)/$(BINARY)
SHA_FILE := $(OUT).sha256

.PHONY: all build release sha clean

all: release sha

build:
	@mkdir -p $(DIST_DIR)
	go build -o $(OUT) .

release:
	@mkdir -p $(DIST_DIR)
	go build -trimpath -ldflags "-s -w" -o $(OUT) .

sha: release
	shasum -a 256 $(OUT) > $(SHA_FILE)
	@echo "Wrote $(SHA_FILE)"

clean:
	rm -rf $(DIST_DIR)
