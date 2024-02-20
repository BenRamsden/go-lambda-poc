package api

import (
	"context"
	"slices"
	"testing"

	"github.com/google/uuid"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/dynamo"
	"github.com/jugo-io/go-poc/api/model"
	"github.com/jugo-io/go-poc/api/service"
)

func LocalDynamoAssetService(id string) func() model.AssetRepository {
	return func() model.AssetRepository {
		return dynamo.NewForTest(id)
	}
}

func LocalDynamoAssetServiceCleanup(repo model.AssetRepository) {
	dynamo.CleanupForTest(repo)
}

func TestCreateAsset(t *testing.T) {
	ctx := context.TODO()

	auth := auth.Auth{
		ID: uuid.New().String(),
	}

	table := []struct {
		name            string
		createAssetRepo func() model.AssetRepository
		cleanUpRepo     func(model.AssetRepository)
	}{
		{
			name:            "CreateAsset",
			createAssetRepo: LocalDynamoAssetService("CreateAsset"),
			cleanUpRepo:     LocalDynamoAssetServiceCleanup,
		},
	}

	assets := []struct {
		name        string
		description string
		uri         string
	}{
		{
			name:        "Test Asset",
			description: "Test Description",
			uri:         "http://test.com",
		},
		{
			name:        "Test Asset 2",
			description: "Test Description 2",
			uri:         "http://test2.com",
		},
		{
			name:        "Test Asset 3",
			description: "Test Description 3",
			uri:         "http://test3.com",
		},
		{
			name:        "Test Asset 4",
			description: "Test Description 4",
			uri:         "http://test4.com",
		},
	}

	for _, tt := range table {
		repo := tt.createAssetRepo()
		service := service.NewAssetService(repo)

		for _, asset := range assets {
			_, err := service.CreateAsset(ctx, auth, model.NewAsset{
				Name:        asset.name,
				Description: asset.description,
				URI:         asset.uri,
			})
			if err != nil {
				t.Errorf("Expected no error, got:\n\t%s", err)
			}
		}

		savedAssets, err := service.GetAssets(ctx, auth)
		if err != nil {
			t.Errorf("Expected no error, got:\n\t%s", err)
		}

		if len(savedAssets) != len(assets) {
			t.Errorf("Expected %d assets, got %d", len(assets), len(savedAssets))
		}

		// Expect assets to be created in order
		slices.SortFunc(savedAssets, func(i, j model.Asset) int {
			if i.CreatedAt.Before(j.CreatedAt) {
				return -1
			}

			if i.CreatedAt.After(j.CreatedAt) {
				return 1
			}

			return 0
		})

		for i, asset := range savedAssets {
			if asset.Name != assets[i].name {
				t.Errorf("Expected asset name to be %s, got %s", assets[i].name, asset.Name)
			}
			if asset.Description != assets[i].description {
				t.Errorf("Expected asset description to be %s, got %s", assets[i].description, asset.Description)
			}
			if asset.URI != assets[i].uri {
				t.Errorf("Expected asset uri to be %s, got %s", assets[i].uri, asset.URI)
			}
		}

		tt.cleanUpRepo(repo)
	}
}
