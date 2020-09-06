package workers

func Assign(params map[string]interface{}) {
	worker := params["Worker"]

	switch {
	case worker == "SendSms":
		SendSmsWorker()
	default:
	}
}
