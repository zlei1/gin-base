package request

type AdminLoginRequest struct {
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
