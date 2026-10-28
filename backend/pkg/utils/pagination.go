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

// instead of writing this normalization everytime, call this method
func (p Pagination) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = DEFAULT_PAGE_SIZE
	}
}

func NewPaginatedData[T any](data []T, total int64, page int, pageSize int) PaginatedData[T] {
	return PaginatedData[T]{
		Data:       data,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
	}
}
