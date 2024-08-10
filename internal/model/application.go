package model

type Response[T any] struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    T      `json:"data"`
}

type PaginationResponse[T any] struct {
	TotalPages int `json:"totalPages"`
	TotalCount int `json:"totalCount"`
	Items      T   `json:"items"`
}

type ErrorResponse struct {
	Response[string]
}

type PaginationRequest struct {
	Search     string `json:"Search"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
}
