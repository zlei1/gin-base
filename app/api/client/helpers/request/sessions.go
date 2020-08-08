package request

type UserLoginRequest struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
	Vcode string `json:"vcode" form:"vcode" validate:"required,len=6"`
}
