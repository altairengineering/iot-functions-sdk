SHELL := /bin/bash
PYTHON_VENV=venv
PYTHON_REQ=requirements.txt
PYTHON_INDEX=index.py
SERVER_PORT=8082

# Color variables
COLOR_GREEN=\033[0;32m
COLOR_RED=\033[0;31m
COLOR_BLUE=\033[0;34m
COLOR_RESET=\033[0m

.PHONY: install run

install:
	@echo -e "🔍 $(COLOR_BLUE)Checking Python version...$(COLOR_RESET)"
	@PYTHON_VERSION=$$(python3 -c 'import sys; print(".".join(map(str, sys.version_info[:3])))'); \
	REQUIRED_VERSION="3.10"; \
 	if [ $$(echo -e "$$PYTHON_VERSION\n$$REQUIRED_VERSION" | sort -V | head -n1) != "$$REQUIRED_VERSION" ]; then \
		echo -e "$(COLOR_RED)Python version must be >= $$REQUIRED_VERSION. Current version: $$PYTHON_VERSION$(COLOR_RESET)"; \
		exit 1; \
	fi

	@echo -e "🛠 $(COLOR_BLUE)Creating Python virtual environment and installing dependencies...$(COLOR_RESET) (python3 -m venv $(PYTHON_VENV))"
	@python3 -m venv $(PYTHON_VENV)

	@echo -e "⚙️ $(COLOR_BLUE)Installing dependencies from requirements.txt...$(COLOR_RESET) ($(PYTHON_VENV)/bin/pip install -r $(PYTHON_REQ))"
	@$(PYTHON_VENV)/bin/pip install -r $(PYTHON_REQ) | sed 's/^/\t/'

	@echo -e "\n🎉 $(COLOR_GREEN)Python virtual environment created at $(PYTHON_VENV)$(COLOR_RESET)"
	@echo -e "\nNext steps:"
	@echo -n "  1. Activate the virtual environment: "
	@echo -e "$(COLOR_GREEN)source venv/bin/activate$(COLOR_RESET)"
	@echo -e "  2. Run the Python server: $(COLOR_GREEN)$(MAKE) run$(COLOR_RESET)"

run:
	@if [ ! -d "$(PYTHON_VENV)" ]; then \
		$(MAKE) install; \
		echo ""; \
	fi
	@echo -e "▶️ $(COLOR_BLUE)Running Python server (port $(SERVER_PORT))...$(COLOR_RESET)"
	export _SECRETS_PATH=../variables; \
	$(PYTHON_VENV)/bin/python $(PYTHON_INDEX)
