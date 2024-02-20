package dynamo

import (
	"testing"
	"time"

	"github.com/jugo-io/go-poc/api/model"
)

func TestAssetFromModel(t *testing.T) {
	model := model.Asset{
		ID:          "id",
		Owner:       "owner",
		Name:        "name",
		Description: "description",
		URI:         "uri",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	asset := AssetFromModel(model)

	// Key Test

	if asset.PK != model.Owner {
		t.Errorf("Expected %s, got %s", model.Owner, asset.PK)
	}

	if asset.SK != model.CreatedAt {
		t.Errorf("Expected %s, got %s", model.CreatedAt, asset.SK)
	}

	// Data Test

	if asset.ID != model.ID {
		t.Errorf("Expected %s, got %s", model.ID, asset.ID)
	}

	if asset.Owner != model.Owner {
		t.Errorf("Expected %s, got %s", model.Owner, asset.Owner)
	}

	if asset.Name != model.Name {
		t.Errorf("Expected %s, got %s", model.Name, asset.Name)
	}

	if asset.Description != model.Description {
		t.Errorf("Expected %s, got %s", model.Description, asset.Description)
	}

	if asset.URI != model.URI {
		t.Errorf("Expected %s, got %s", model.URI, asset.URI)
	}

	if asset.CreatedAt != model.CreatedAt {
		t.Errorf("Expected %s, got %s", model.CreatedAt, asset.CreatedAt)
	}

	if asset.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected %s, got %s", model.UpdatedAt, asset.UpdatedAt)
	}
}
