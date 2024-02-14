// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Asset struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	URI         string `json:"URI"`
	CreatedAt   string `json:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt"`
}

type Header struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Me struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Mutation struct {
}

type NewAsset struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type NewAssetResponse struct {
	Asset        Asset        `json:"Asset"`
	PresignedURL PresignedURL `json:"PresignedURL"`
}

type Pagination struct {
	Total  int     `json:"total"`
	Limit  int     `json:"limit"`
	Cursor *string `json:"cursor,omitempty"`
}

type PresignedURL struct {
	URL    string   `json:"URL"`
	Fields []Header `json:"Fields"`
}

type Query struct {
}

type UpdateAsset struct {
	ID          string  `json:"ID"`
	Name        *string `json:"Name,omitempty"`
	Description *string `json:"Description,omitempty"`
}