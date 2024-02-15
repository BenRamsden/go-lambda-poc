package dynamo

import (
	"time"

	"github.com/jugo-io/go-poc/api/model"
)

type Asset struct {
	PK string // Owner ID
	SK time.Time

	ID          string    `dynamodbav:"id" json:"id"`
	Owner       string    `dynamodbav:"owner" json:"owner"`
	Name        string    `dynamodbav:"name" json:"name"`
	Description string    `dynamodbav:"description" json:"description"`
	URI         string    `dynamodbav:"uri" json:"uri"`
	CreatedAt   time.Time `dynamodbav:"created_at" json:"created_at"`
	UpdatedAt   time.Time `dynamodbav:"updated_at" json:"updated_at"`
}

func AssetFromModel(model model.Asset) Asset {
	return Asset{
		PK: model.Owner,
		SK: model.CreatedAt,

		ID:          model.ID,
		Owner:       model.Owner,
		Name:        model.Name,
		Description: model.Description,
		URI:         model.URI,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func ModelFromAsset(asset Asset) model.Asset {
	return model.Asset{
		ID:          asset.ID,
		Owner:       asset.Owner,
		Name:        asset.Name,
		Description: asset.Description,
		URI:         asset.URI,
		CreatedAt:   asset.CreatedAt,
		UpdatedAt:   asset.UpdatedAt,
	}
}
