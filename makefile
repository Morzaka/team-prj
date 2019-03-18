include env

VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")


GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

LINTERCOMMAND=golangci-lint  run




packages = \
	./configurations \
	./logger \
	./services/authorization\models\
	./services/authorization\session\
        ./services\database\
	./services\handlers\
	./train\
	
       
setup:
	go get -u golang.org/x/vgo
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint


go-compile:  vgo get -u && vgo build

build: 
        vgo build
run:	
        vgo run ./main.go
vendor:	
	vgo vendor
lint:
	golangci-lint  run ./src/...


.PHONY: all build vendor test lint

all:
	lint test

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY: test

	@$(foreach package,$(packages), \
		set -e; \
		vgo test -all

.PHONY: code-quality-print
code-quality-print:
	$(LINTERCOMMAND) ./...


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

dc-build:
	docker-compose build
dc-up:
	docker-compose up &

