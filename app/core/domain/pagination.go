package domain

type Pagination[T any] struct {
	Objects 		T `json:"objects"`
	Total 			int32 `json:"total"`
}