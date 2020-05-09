#!/bin/bash

PROJECTNAME=$(shell basename "$(PWD)")
BINARY=engine
FILE=./bin/golangci-lint

# include .env

check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

## dependency-prepare: download dep package manager in order to run make dependency
setup-dependency:
	@echo "  > Installing dep package manager ..."
	@- go get -d -u github.com/golang/dep

## dependency: Ensure all packages are installed
dependency:
	@echo "  > Performing go package installation ..."
	@- dep ensure

## coverage: run test coverage
test-coverage:
	@echo "  >  Performing Code Test coverage ..."
	@- go test -v -cover ./...

## build: Clean build code as binary
build:
	@echo "  >  Performing Code Build"
	@- go build -o ${BINARY} main.go

## unittest: Run unit testing for code
test:
	@echo "  >  Performing Code Test ..."
	@- go test -short  ./...

## lint: Run code linting and install golangci-lint in case it isn't present on folder
lint:
ifneq ("$(wildcard $(FILE))","")
    
	@echo "  >  Executing linting of current code ..."
	@- ./bin/golangci-lint run --exclude-use-default=false --enable=golint --enable=gocyclo --enable=goconst --enable=unconvert ./...
else
    
	@echo "  > lint-prepare not present"

	@echo "  > downloading golangci-lint package"

	@- curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
endif

## docker-setup: Runs MySQL and PHPMyAdmin Docker containers
docker-mysql:
	@echo "  > Installing MySQL docker container"
	@- docker run -p 3306:3306 --name muble-mysql --env-file .env -d mysql:latest
	
docker-phpmyadmin:
	@echo "  > Installing MySQL Web UI (PHPMyAdmin)"
	@- docker run --name muble-phpmyadmin -d --link muble-mysql:db -p 8080:80 -v ${HOME}/config.user.inc.php:/etc/phpmyadmin/config.user.inc.php phpmyadmin/phpmyadmin
	@echo "  > phpmyadmin config file located at:  $(HOME}/config.user.inc.php) "



## setup-envs: Setup Enviroment variables
setup-env:
	@- if test -f .env; \
	then export $(cat .env); \
	fi

## run: Run go server
run-server:
	@echo "  >  Running server ..."
	go run main.go

## clean: Remove go build binary from folder
clean:
	@echo "  > Cleaning binary generated code ..."
	@- if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo