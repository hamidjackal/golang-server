SHELL         := /bin/bash
.DEFAULT_GOAL := help
APP_NAME      := golang-server
ENV           := dev

run/watch: ## Runs the server in watch mode
	@echo "Start running the server in watch mode"
	cd server && rm ./tmp/main && go build -o ./tmp/main && ./tmp/main
	

run/build: ## Builds the server
	@echo "Start running the server in watch mode"
	cd server && go build -o ./tmp/main