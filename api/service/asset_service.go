package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/model"
)

type assetService struct {
	repo model.AssetRepository
}

// CreateAsset implements model.AssetService.
func (svc *assetService) CreateAsset(auth auth.Auth, newAsset model.NewAsset) (model.Asset, error) {
	return svc.repo.CreateAsset(model.Asset{
		ID:          uuid.NewString(),
		Owner:       auth.ID,
		Name:        newAsset.Name,
		Description: newAsset.Description,
		URI:         newAsset.URI,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}

// GetAssets implements model.AssetService.
func (svc *assetService) GetAssets(auth auth.Auth) ([]model.Asset, error) {
	return svc.repo.GetAssets(auth.ID)
}

func NewAssetService(repo model.AssetRepository) model.AssetService {
	return &assetService{repo: repo}
}
