package response

type CaptchaResponse struct {
	CaptchaToken string `json:"captcha_token"`
	Captcha      string `json:"captcha"`
}
