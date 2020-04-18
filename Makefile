DEV_IMAGE := drawbridge
DOCKER_COMPOSE ?= $(shell which docker-compose)
DOCKER_COMPOSE_DEV_RUN := $(DOCKER_COMPOSE) run --rm $(DEV_IMAGE)

.SILENT: help
help: ## Show this help message
	echo "usage: make [target] ..."
	echo ""
	echo "targets:"
	fgrep --no-filename "##" $(MAKEFILE_LIST) | fgrep --invert-match $$'\t' | sed -e 's/: ## / - /'

build-image: ## Build the docker image
	$(DOCKER_COMPOSE) build

update: ## Update dependencies
	rm -rf vendor
	$(DOCKER_COMPOSE_DEV_RUN) dep ensure

dep-add: ## Add a new dependency to Gopkg.lock
	$(DOCKER_COMPOSE_DEV_RUN) dep ensure -add ${pkg}

build: ## Build and compile the source
	$(DOCKER_COMPOSE_DEV_RUN) go build -o bin/drawbridge # This will go into /go/src/drawbridge/bin

start: ## Start the dev service
	$(DOCKER_COMPOSE) run --service-ports --rm $(DEV_IMAGE) /go/bin/drawbridge /config/config.yaml

test: ## Run the tests
	$(DOCKER_COMPOSE_DEV_RUN) go test

.PHONY: clean
clean: ## Clean up any containers and images
	rm -r vendor/
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) rm -f
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans
