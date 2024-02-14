package model

type PaginationDirection string

type Pagination struct {
	Count int `json:"count"`
	Page  int `json:"page"`
}
