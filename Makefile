export ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

PROJECT_NAME := goodtools

GIT_COMMIT   := $(shell git rev-parse --short HEAD)
GIT_DIRTY    := $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
VERSION      ?= $(GIT_COMMIT)

DOCKER_SERVICE   := dev
DOCKER_COMPOSE   := $(shell which docker-compose)
DOCKER_SHELL     := $(DOCKER_COMPOSE) exec -T $(DOCKER_SERVICE) /bin/bash
DOCKER_SHELL_TTY := $(DOCKER_COMPOSE) exec $(DOCKER_SERVICE) /bin/bash

.PHONY: help
help:
	@echo "help:     This help message"
	@echo "info:     Print environment information"
	@echo "up:       Create development environment"
	@echo "down:     Destroy development environment"

.PHONY: info
info:
	@echo "Git Commit:      $(GIT_COMMIT)"
	@echo "Git Tree State:  $(GIT_DIRTY)"

.PHONY: up
up:
	$(DOCKER_COMPOSE) up --build -d

.PHONY: down
down:
	-$(DOCKER_COMPOSE) down

.PHONY: shell
shell: SHELL := $(DOCKER_SHELL_TTY)
shell:
	bash

.PHONY: deps
deps: SHELL := $(DOCKER_SHELL)
deps: ./vendor

.vendor: SHELL := $(DOCKER_SHELL)
./vendor:
	glide install -v

create: SHELL := $(DOCKER_SHELL)
create: deps
	go fmt ./...
	go build -o bin/poc cmd/poc/main.go

.PHONY: test-unit
test-unit: SHELL := $(DOCKER_SHELL)
test-unit:
	go test $$(glide novendor)
