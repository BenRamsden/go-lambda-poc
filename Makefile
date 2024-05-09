.PHONY: init api ui package generate deploy

MAKEFLAGS += --silent

init:
	@echo "Installing Go dependencies"
	@go mod tidy
	@cd pulumi && go mod tidy

	@echo "Installing TypeScript dependencies"
	@cd graphql/typescript && yarn install
	@cd ui && yarn install

	@echo "Installing tmuxinator"
	@command -v gem >/dev/null 2>&1 || { echo "ruby is not installed. Please install RubyGems to continue."; exit 1; }
	@gem install tmuxinator

start:
	@command -v tmux >/dev/null 2>&1 || { echo "tmux is not installed. Please install tmux to continue."; exit 1; }
	@command -v tmuxinator >/dev/null 2>&1 || { echo "tmuxinator is not installed. Please install tmuxinator to continue. You may also need to add the gem bin directory to your PATH."; exit 1; }
	@tmuxinator start

stop:
	@tmuxinator stop go-lambda-poc

ui:
	cd ui && yarn dev

api:
	go build -ldflags="-s -w"  -o bin/api ./cmd/api/local_main.go
	./bin/api

db:
	docker-compose up

generate:
	@echo "Generating Go code"
	cd graphql/go && go run github.com/99designs/gqlgen generate --config api-gqlgen.yml
	@echo "Generating TypeScript code"
	cd graphql/typescript && yarn generate

LAMBDAS_BINARIES := api
package:
	@echo "Building lambdas"
	@for lambda in $(LAMBDAS_BINARIES); do \
		GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o "bin/lambda/$$lambda/bootstrap" "cmd/$$lambda/lambda_main.go"; \
		zip -j "bin/lambda/$$lambda/$$lambda.zip" "bin/lambda/$$lambda/bootstrap"; \
	done
	@echo "Building UI"
	@cd ui && yarn package --mode sandbox

preview:
	cd pulumi && pulumi preview

deploy:
	cd pulumi && pulumi up -y

test:
	docker-compose up -d --build
	go clean -testcache
	go test ./...
	docker-compose down --remove-orphans