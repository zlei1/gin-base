package response

type PaginatorResponse struct {
	Items      interface{} `json:"items"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
}
