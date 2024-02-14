.PHONY: list local package generate

MAKEFLAGS += --silent

init:
	@echo "Installing Go dependencies"
	@go mod tidy
	@cd pulumi && go mod tidy
	@echo "Installing TypeScript dependencies"
	@cd graphql/typescript && yarn install
	@cd ui && yarn install

api:
	go run cmd/api/local_main.go

generate:
	@echo "Generating Go code"
	cd graphql/go && go run github.com/99designs/gqlgen generate --config api-gqlgen.yml
	@echo "Generating TypeScript code"
	cd graphql/typescript && yarn generate

LAMBDAS_BINARIES := api migrate
package:
	@echo "Building lambdas"
	@for lambda in $(LAMBDAS_BINARIES); do \
		GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o "bin/lambda/$$lambda/bootstrap" "cmd/$$lambda/lambda_main.go"; \
		zip -j "bin/lambda/$$lambda/$$lambda.zip" "bin/lambda/$$lambda/bootstrap"; \
	done
	@echo "Building UI"
	@cd ui && yarn package
	