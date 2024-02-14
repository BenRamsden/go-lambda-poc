package model

import (
	"context"
	"time"
)

type NewAsset struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateAsset struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type Asset struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URI         string    `json:"uri"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PresignedURL struct {
	URL    string            `json:"url"`
	Fields map[string]string `json:"fields"`
}

type AssetRepository interface {
	CreateAsset(context context.Context, newAsset NewAsset) (Asset, error)
	UpdateAsset(context context.Context, updateAsset UpdateAsset) (Asset, error)
	UpdateAssetURI(context context.Context, id string, uri string) error
	DeleteAsset(context context.Context, id string) error

	GetAsset(context context.Context, id string) (Asset, error)
	GetAssets(context context.Context, pagination Pagination) ([]Asset, error)
}

type AssetService interface {
	CreateAsset(context context.Context, newAsset NewAsset) (Asset, error)
	UpdateAsset(context context.Context, updateAsset UpdateAsset) (Asset, error)
	DeleteAsset(context context.Context, id string) error
	UploadAsset(context context.Context, id string) (PresignedURL, error)

	GetAsset(context context.Context, id string) (Asset, error)
	GetAssets(context context.Context, pagination Pagination) ([]Asset, error)
}
