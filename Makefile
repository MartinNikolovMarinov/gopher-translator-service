PWD := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
CMD_DIR ?= $(PWD)cmd
OUT_DIR ?= $(PWD)out
BIN_NAME ?= main

## HELPERS
.PHONY: help
help:
	@echo Supported targets:
	@cat $(MAKEFILE_LIST) | grep -e "^[\.a-zA-Z0-9_-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-35s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

## TARGETS
.PHONY: build
build: ## compile executables
	mkdir -p $(OUT_DIR)
	cd $(CMD_DIR) && go build -o $(OUT_DIR)/$(BIN_NAME) && cd $(PWD)

.PHONY: run_main
run_main: ## run main
	go run $(CMD_DIR)/main.go

.PHONY: test
test: clean ## run tests
	go test ./... -v -cover -coverprofile=out/cover.txt -bench=. && \
	go tool cover -html=out/cover.txt -o out/cover.html

.PHONY: clean
clean: ## clean output folder
	rm -rf $(OUT_DIR)/*
