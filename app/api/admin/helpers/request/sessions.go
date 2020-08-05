package request

type PhoneLoginRequest struct {
	Phone        string `json:"phone"`
	VerifyCode   string `json:"verify_code"`
	Captcha      string `json:"captcha"`
	CaptchaToken string `json:"captcha_token"`
}
