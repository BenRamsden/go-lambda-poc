package sql

import (
	"context"

	"github.com/jugo-io/go-poc/internal/api/model"
)

// CreateAsset implements Repository.
func (*repository) CreateAsset(context context.Context, newAsset model.NewAsset) (model.Asset, error) {
	panic("unimplemented")
}

// DeleteAsset implements Repository.
func (*repository) DeleteAsset(context context.Context, id string) error {
	panic("unimplemented")
}

// GetAsset implements Repository.
func (*repository) GetAsset(context context.Context, id string) (model.Asset, error) {
	panic("unimplemented")
}

// GetAssets implements Repository.
func (*repository) GetAssets(context context.Context, pagnation model.Pagination) ([]model.Asset, error) {
	panic("unimplemented")
}

// UpdateAsset implements Repository.
func (*repository) UpdateAsset(context context.Context, updateAsset model.UpdateAsset) (model.Asset, error) {
	panic("unimplemented")
}

// UpdateAssetURI implements Repository.
func (*repository) UpdateAssetURI(context context.Context, id string, uri string) error {
	panic("unimplemented")
}
