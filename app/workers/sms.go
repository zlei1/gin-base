package workers

import (
	"log"
)

type SendSmsParams struct {
	Worker string
	Phone  string
	Vcode  string
}

func SendSmsWorker() {
	log.Println("Send Sms Worker Succeed")
}
