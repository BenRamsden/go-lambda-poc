package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/model"
)

type assetService struct {
	repo model.AssetRepository
}

// CreateAsset implements model.AssetService.
func (svc *assetService) CreateAsset(ctx context.Context, auth auth.Auth, newAsset model.NewAsset) (model.Asset, error) {
	return svc.repo.CreateAsset(ctx, model.Asset{
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
func (svc *assetService) GetAssets(ctx context.Context, auth auth.Auth) ([]model.Asset, error) {
	return svc.repo.GetAssets(ctx, auth.ID)
}

func NewAssetService(repo model.AssetRepository) model.AssetService {
	return &assetService{repo: repo}
}
