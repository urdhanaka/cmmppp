package utils

const DEFAULT_PAGE_SIZE = 10

type (
	// Pagination is 1-indexed; Page/PageSize below 1 fall back to defaults.
	Pagination struct {
		Page     int
		PageSize int
	}

	PaginatedData[T any] struct {
		Data       []T   `json:"data"`
		TotalCount int64 `json:"total_count"`
		Page       int   `json:"page"`
		PageSize   int   `json:"page_size"`
	}
)

func NewPaginatedData[T any](data []T, total int64, page int, pageSize int) PaginatedData[T] {
	return PaginatedData[T]{
		Data:       data,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
	}
}
