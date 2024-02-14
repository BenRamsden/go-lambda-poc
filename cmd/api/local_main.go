//go:build !lambda

package main

import "github.com/jugo-io/go-poc/internal/api"

func main() {
	r := api.Handler(api.HandlerOptions{})
	r.Run(":8080")
}
