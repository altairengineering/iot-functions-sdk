SHELL := /bin/bash
GO_MOD=go.mod
SERVER_PORT=8082

# Color variables
COLOR_GREEN=\033[0;32m
COLOR_RED=\033[0;31m
COLOR_BLUE=\033[0;34m
COLOR_RESET=\033[0m

.PHONY: install run

install:
	@echo -e "🔍 $(COLOR_BLUE)Checking Go version...$(COLOR_RESET)"
	@GO_VERSION=$$(go version | awk '{print $$3}' | sed 's/go//'); \
	REQUIRED_VERSION="1.19"; \
	if [ $$(echo -e "$$GO_VERSION\n$$REQUIRED_VERSION" | sort -V | head -n1) != "$$REQUIRED_VERSION" ]; then \
		echo -e "$(COLOR_RED)Go version must be >= $$REQUIRED_VERSION. Current version: $$GO_VERSION$(COLOR_RESET)"; \
		exit 1; \
	fi

	@echo -e "⚙️ $(COLOR_BLUE)Installing Go dependencies...$(COLOR_RESET) (go mod tidy)"
	@go mod tidy | sed 's/^/\t/'

	@echo -e "\n🎉 $(COLOR_GREEN)Go dependencies installed$(COLOR_RESET)"
	@echo -e "\nNext steps:"
	@echo -e "  1. Run the Go server: $(COLOR_GREEN)$(MAKE) run$(COLOR_RESET)"

run:
	@echo -e "▶️ $(COLOR_BLUE)Running Go server (port $(SERVER_PORT))...$(COLOR_RESET)"
	@export _SECRETS_PATH=../variables; \
	go run .
