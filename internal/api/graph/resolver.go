package graph

import "github.com/jugo-io/go-poc/internal/api/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AssetService model.AssetService
}
