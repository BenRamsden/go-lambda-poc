.PHONY: list local package generate

MAKEFLAGS += --silent

api:
	go run cmd/api/local_main.go

generate:
	cd internal/api/graph && go run github.com/99designs/gqlgen generate


LAMBDAS_BINARIES := api migrate
package:
	@echo "Building lambdas"
	@for lambda in $(LAMBDAS_BINARIES); do \
		GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o "bin/lambda/$$lambda/bootstrap" "cmd/$$lambda/lambda_main.go"; \
		zip -j "bin/lambda/$$lambda/$$lambda.zip" "bin/lambda/$$lambda/bootstrap"; \
	done
	@echo "Building UI"
	@cd ui && yarn package
	