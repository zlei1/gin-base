package request

type AdminLoginRequest struct {
	Phone    string `json:"phone" form:"phone" binding:"required`
	Password string `json:"password" form:"password" binding:"required`
}
