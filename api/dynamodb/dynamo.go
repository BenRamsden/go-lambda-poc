package dynamodb

import (
	"github.com/jugo-io/go-poc/api/model"
	"github.com/jugo-io/go-poc/api/service"
)

type Repository interface {
	service.AssetRepository
}

type repository struct {
}

func New() Repository {
	return &repository{}
}

// CreateAsset implements Repository.
func (*repository) CreateAsset(ownerId string, asset model.NewAsset) (model.Asset, error) {
	panic("unimplemented")
}

// GetAssets implements Repository.
func (*repository) GetAssets(ownerId string) ([]model.Asset, error) {
	panic("unimplemented")
}
