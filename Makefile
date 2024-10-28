.PHONY: run-python run-go install-python install-go

PYTHON_DIR=python3-http
GO_DIR=golang-http

run-python:
	@$(MAKE) -C $(PYTHON_DIR) run

run-go:
	@$(MAKE) -C $(GO_DIR) run

install-python:
	@$(MAKE) -C $(PYTHON_DIR) install

install-go:
	@$(MAKE) -C $(GO_DIR) install
