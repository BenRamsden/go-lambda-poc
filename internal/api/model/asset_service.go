package model

import (
	"context"
	"errors"
)

type assetService struct {
	repo AssetRepository
}

// CreateAsset implements AssetService.
func (a *assetService) CreateAsset(context context.Context, newAsset NewAsset) (Asset, error) {
	if len(newAsset.Name) == 0 {
		return Asset{}, NewInvalidInputError(errors.New("name is required"))
	}

	return a.repo.CreateAsset(context, newAsset)
}

// DeleteAsset implements AssetService.
func (a *assetService) DeleteAsset(context context.Context, id string) error {
	if len(id) == 0 {
		return NewInvalidInputError(errors.New("id is required"))
	}

	return a.repo.DeleteAsset(context, id)
}

// GetAsset implements AssetService.
func (a *assetService) GetAsset(context context.Context, id string) (Asset, error) {
	if len(id) == 0 {
		return Asset{}, NewInvalidInputError(errors.New("id is required"))
	}

	return a.repo.GetAsset(context, id)
}

// GetAssets implements AssetService.
func (a *assetService) GetAssets(context context.Context, pagination Pagination) ([]Asset, error) {
	return a.repo.GetAssets(context, pagination)
}

// UpdateAsset implements AssetService.
func (a *assetService) UpdateAsset(context context.Context, updateAsset UpdateAsset) (Asset, error) {
	if len(updateAsset.ID) == 0 {
		return Asset{}, NewInvalidInputError(errors.New("id is required"))
	}

	return a.repo.UpdateAsset(context, updateAsset)
}

// UploadAsset implements AssetService.
func (a *assetService) UploadAsset(context context.Context, id string) (PresignedURL, error) {
	panic("unimplemented")
}

func NewAssetService(repo AssetRepository) AssetService {
	return &assetService{repo: repo}
}
