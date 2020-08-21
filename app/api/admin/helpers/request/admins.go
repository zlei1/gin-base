package request

type IndexAdminRequest struct {
	PaginatorRequest
}

type CreateAdminRequest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Phone string `json:"phone" form:"phone" validate:"required"`
}
