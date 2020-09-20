package anticaptcha

type ReportIncorrectRecaptcha_s struct {
	response_s
	Status string `json:"status"`
}

func ReportIncorrectRecaptcha (clientKey string, taskId int) (*ReportIncorrectRecaptcha_s, error) {
	resp_s := &ReportIncorrectRecaptcha_s{}
	err := getApiContent("reportIncorrectRecaptcha ", Obj{
		"clientKey": clientKey,
		"taskId": taskId,
	}, resp_s)
	if err != nil {
		return nil, err
	}
	return resp_s, nil
}