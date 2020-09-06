package request

type PhoneVerifyCodeRequest struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
}
