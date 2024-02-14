//go:build lambda

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/jugo-io/go-poc/internal/api"
)

func main() {
	adapter := ginadapter.NewV2(api.Handler(api.HandlerOptions{}))
	lambda.Start(adapter.ProxyWithContext)
}
