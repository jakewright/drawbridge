DOCKER_COMPOSE ?= $(shell which docker-compose)
DOCKER_COMPOSE_RUN := $(DOCKER_COMPOSE) run --rm drawbridge

.SILENT: help
help: ## Show this help message
	echo "usage: make [target] ..."
	echo ""
	echo "targets:"
	fgrep --no-filename "##" $(MAKEFILE_LIST) | fgrep --invert-match $$'\t' | sed -e 's/: ## / - /'

build-image: ## Build the docker image
	$(DOCKER_COMPOSE) build

install: build-image ## Install dependencies
	$(DOCKER_COMPOSE) build
	$(DOCKER_COMPOSE_RUN) glide install
	$(DOCKER_COMPOSE_RUN) go install

update: ## Update dependencies
	rm -rf vendor
	$(DOCKER_COMPOSE_RUN) glide update

glide-get: ## Add a new dependency to glide.yaml
	$(DOCKER_COMPOSE_RUN) glide get ${pkg}

build: ## Build and compile the source
	$(DOCKER_COMPOSE_RUN) go build -o bin/drawbridge # This will go into /go/src/drawbridge/bin

start: ## Start the service
	$(DOCKER_COMPOSE) up

test: ## Run the tests
	$(DOCKER_COMPOSE_RUN) go test

.PHONY: clean
clean: ## Clean up any containers and images
	rm -r vendor/
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) rm -f
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans
