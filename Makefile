.PHONY: list local package

LAMBDAS_BINARIES := $(shell find ./cmd -name "lambda_main.go" | xargs -n1 dirname | xargs -n1 basename)
LOCAL_BINARIES := $(shell find ./cmd -name "local_main.go" | xargs -n1 dirname | xargs -n1 basename)

list:
	@echo "Available Local Targets:"
	@echo $(LOCAL_BINARIES) | xargs -n1 echo "-"
	@echo "Available Lambda Targets:"
	@echo $(LAMBDAS_BINARIES) | xargs -n1 echo "-"

ifeq (local,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_LOCAL_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_LOCAL_ARGS):;@:)
endif

local:
	@go build -o bin/local/$(RUN_LOCAL_ARGS) cmd/$(RUN_LOCAL_ARGS)/local_main.go
	@./bin/local/$(RUN_LOCAL_ARGS)

package:
	@echo $(LAMBDAS_BINARIES) | GOARCH=arm64 GOOS=linux xargs -I{} go build -ldflags="-s -w" -o "bin/lambda/{}/bootstrap" "cmd/{}/lambda_main.go"
	@echo $(LAMBDAS_BINARIES) | xargs -I{} zip -j "bin/lambda/{}/{}.zip" "bin/lambda/{}/bootstrap"
	
generate:
	cd internal/api/graph && go run github.com/99designs/gqlgen generate