package request

type AdminLoginRequest struct {
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	Captcha      string `json:"captcha"`
	CaptchaToken string `json:"captcha_token"`
}
