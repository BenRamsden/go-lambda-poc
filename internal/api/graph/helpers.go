package graph

import "github.com/jugo-io/go-poc/internal/api/model"

func ConvertAsset(a model.Asset) Asset {
	return Asset{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		URI:         a.URI,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func P(pagination *Pagination) model.Pagination {
	if pagination == nil {
		return model.Pagination{
			Count: 10,
			Page:  0,
		}
	}

	return model.Pagination{
		Count: pagination.Count,
		Page:  pagination.Page,
	}
}
