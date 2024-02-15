//go:build !lambda

package main

import (
	"github.com/jugo-io/go-poc/api"
)

func main() {
	options := api.HandlerOptions{}

	r := api.Handler(options)
	r.Run(":4000")
}
