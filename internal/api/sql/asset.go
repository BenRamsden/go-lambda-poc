package sql

import (
	"context"

	"github.com/jugo-io/go-poc/internal/api/model"
)

type Asset struct {
	UUIDModel
	OwnerId     string `gorm:"index"`
	Name        string
	Description string
	URI         string
}

// CreateAsset implements Repository.
func (r *repository) CreateAsset(context context.Context, newAsset model.NewAsset) (model.Asset, error) {
	asset := &Asset{
		OwnerId:     newAsset.OwnerId,
		Name:        newAsset.Name,
		Description: newAsset.Description,
		URI:         "",
	}

	tx := r.Create(asset)

	if tx.Error != nil {
		return model.Asset{}, tx.Error
	}

	return model.Asset{
		ID:          asset.ID,
		OwnerId:     asset.OwnerId,
		Name:        asset.Name,
		Description: asset.Description,
		URI:         asset.URI,
		CreatedAt:   asset.CreatedAt,
		UpdatedAt:   asset.UpdatedAt,
	}, nil
}

// DeleteAsset implements Repository.
func (r *repository) DeleteAsset(context context.Context, id string) error {
	tx := r.Delete(&Asset{
		UUIDModel: UUIDModel{ID: id},
	})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetAsset implements Repository.
func (r *repository) GetAsset(context context.Context, id string) (model.Asset, error) {
	var asset Asset
	tx := r.First(&asset, id)

	if tx.Error != nil {
		return model.Asset{}, tx.Error
	}

	return model.Asset{
		ID:          asset.ID,
		OwnerId:     asset.OwnerId,
		Name:        asset.Name,
		Description: asset.Description,
		URI:         asset.URI,
		CreatedAt:   asset.CreatedAt,
		UpdatedAt:   asset.UpdatedAt,
	}, nil
}

// GetAssets implements Repository.
func (r *repository) GetAssets(context context.Context, filter model.GetAssetsFilter, pagination model.Pagination) ([]model.Asset, error) {
	var assets []Asset

	tx := r.Scopes(P(pagination)).Where("OwnerId = ?", filter.OwnerId).Find(&assets)

	if tx.Error != nil {
		return nil, tx.Error
	}

	result := make([]model.Asset, 0, len(assets))
	for _, asset := range assets {
		result = append(result, model.Asset{
			ID:          asset.ID,
			Name:        asset.Name,
			Description: asset.Description,
			URI:         asset.URI,
			CreatedAt:   asset.CreatedAt,
			UpdatedAt:   asset.UpdatedAt,
		})
	}

	return result, nil
}

// UpdateAsset implements Repository.
func (r *repository) UpdateAsset(context context.Context, updateAsset model.UpdateAsset) (model.Asset, error) {
	asset := &Asset{
		UUIDModel: UUIDModel{ID: updateAsset.ID},
	}

	// Could probably be optimized to happen in a single query
	tx := r.First(asset)

	if tx.Error != nil {
		return model.Asset{}, tx.Error
	}

	if updateAsset.Name != nil {
		asset.Name = *updateAsset.Name
	}

	if updateAsset.Description != nil {
		asset.Description = *updateAsset.Description
	}

	tx = r.Save(asset)

	if tx.Error != nil {
		return model.Asset{}, tx.Error
	}

	return model.Asset{
		ID:          asset.ID,
		Name:        asset.Name,
		Description: asset.Description,
		URI:         asset.URI,
		CreatedAt:   asset.CreatedAt,
		UpdatedAt:   asset.UpdatedAt,
	}, nil
}

// UpdateAssetURI implements Repository.
func (r *repository) UpdateAssetURI(context context.Context, id string, uri string) error {
	tx := r.Model(&Asset{UUIDModel: UUIDModel{ID: id}}).Update("uri", uri)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
