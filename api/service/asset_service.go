package service

import (
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/model"
)

type AssetRepository interface {
	CreateAsset(ownerId string, asset model.NewAsset) (model.Asset, error)
	GetAssets(ownerId string) ([]model.Asset, error)
}

type assetService struct {
	repo AssetRepository
}

// CreateAsset implements model.AssetService.
func (*assetService) CreateAsset(auth auth.Auth, newAsset model.NewAsset) (model.Asset, error) {
	panic("unimplemented")
}

// GetAssets implements model.AssetService.
func (*assetService) GetAssets(auth auth.Auth) ([]model.Asset, error) {
	panic("unimplemented")
}

func NewAssetService(repo AssetRepository) model.AssetService {
	return &assetService{repo: repo}
}
