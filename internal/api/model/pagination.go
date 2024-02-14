package model

type PaginationDirection string

const (
	PaginationDirection__Next PaginationDirection = "next"
	PaginationDirection__Prev PaginationDirection = "prev"
)

type Pagination struct {
	Count  int                 `json:"count"`
	Cursor string              `json:"cursor"`
	Dir    PaginationDirection `json:"dir"`
}
