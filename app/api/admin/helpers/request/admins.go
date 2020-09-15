package request

type IndexAdminRequest struct {
	PaginatorRequest
}

type AdminRequest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Phone string `json:"phone" form:"phone" validate:"required"`
}
