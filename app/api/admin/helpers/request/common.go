package request

type PaginatorRequest struct {
	Page    int `json:"page" form:"page"`
	PerPage int `json:"per_page" form:"per_page"`
}
