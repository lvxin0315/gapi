package tools

type Pagination struct {
	Total    uint                   `json:"total"`
	Page     uint                   `json:"page"`
	PageSize uint                   `json:"page_size"`
	Where    map[string]interface{} `json:"where"`
}

func NewPagination(total, page, pageSize uint, where map[string]interface{}) (pagination Pagination, err error) {
	pagination.Total = total
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Where = where
	return
}
