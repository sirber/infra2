.PHONY: all build run clean help
.DEFAULT_GOAL := help

OUTPUT=build/infra

# desc: Show this help
help:
	@awk '/^# desc:/ {desc=$$0; getline; if ($$0 ~ /^([a-zA-Z0-9_-]+):/) { \
		sub(/^# desc: /, "", desc); \
		sub(":.*$$", "", $$0); \
		printf "  %-8s - %s\n", $$0, desc; \
	}}' $(MAKEFILE_LIST) | ( \
		echo "Docker Self-Hosted Infra"; \
		echo; \
		echo "Usage: make <target> [ACTION=<action>]"; \
		echo "Available targets:"; \
		cat; \
	)

# desc: Build for Linux (output: build/infra)
build:
	@echo "Building for Linux (GOOS=linux, GOARCH=amd64)..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT) infra.go
	@command -v upx >/dev/null 2>&1 && upx -9 -q $(OUTPUT) >/dev/null 2>&1

# desc: Run with ACTION=<action> (up, down, pull, backup, restart)
run:
	@go run . $$ACTION

# desc: Remove build output
clean:
	rm -f $(OUTPUT)