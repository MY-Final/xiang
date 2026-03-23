package dto

type PaginationQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func (q *PaginationQuery) Normalize() {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PageSize <= 0 {
		q.PageSize = 10
	}
	if q.PageSize > 100 {
		q.PageSize = 100
	}
}
