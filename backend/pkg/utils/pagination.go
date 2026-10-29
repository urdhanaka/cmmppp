package utils

const DEFAULT_PAGE_SIZE = 10

type (
	// PaginationMeta is 1-indexed; Page/PageSize below 1 fall back to defaults.
	PaginationMeta struct {
		TotalCount int64 `json:"total_count"`
		Page       int   `json:"page"`
		PageSize   int   `json:"page_size"`
		TotalPages int   `json:"total_pages"`
	}

	PaginatedData[T any] struct {
		Items []T            `json:"items"`
		Meta  PaginationMeta `json:"meta"`
	}
)

// instead of writing this normalization everytime, call this method
func (p *PaginationMeta) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = DEFAULT_PAGE_SIZE
	}
}

func NewPaginatedData[T any](data []T, total int64, page int, pageSize int) PaginatedData[T] {
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	if totalPages < 1 {
		totalPages = 1
	}

	return PaginatedData[T]{
		Items: data,
		Meta: PaginationMeta{
			TotalCount: total,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		},
	}
}
