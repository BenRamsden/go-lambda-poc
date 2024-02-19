package model

import (
	"context"
	"time"

	"github.com/jugo-io/go-poc/api/auth"
)

type NewAsset struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URI         string `json:"uri"`
}

type Asset struct {
	ID          string    `json:"id"`
	Owner       string    `json:"owner"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URI         string    `json:"uri"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AssetRepository interface {
	CreateAsset(ctx context.Context, asset Asset) (Asset, error)
	GetAssets(ctx context.Context, owner string) ([]Asset, error)
}

type AssetService interface {
	CreateAsset(ctx context.Context, auth auth.Auth, newAsset NewAsset) (Asset, error)
	GetAssets(ctx context.Context, auth auth.Auth) ([]Asset, error)
}
