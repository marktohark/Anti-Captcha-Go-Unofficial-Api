package anticaptcha

type Adv_RecaptchaV3Task struct {
	clientKey string
	websiteUrl string
	websiteKey string
	minScore float64
	pageAction string
}

func (this *Adv_RecaptchaV3Task)New(_clientKey string, _websiteUrl string, _websiteKey string, _minScore float64, pageAction string) *Adv_RecaptchaV3Task {
	this.clientKey = _clientKey
	this.websiteUrl = _websiteUrl
	this.websiteKey = _websiteKey
	this.minScore = _minScore
	this.pageAction = pageAction
	return this
}

func (this *Adv_RecaptchaV3Task)Do() (*GetTaskResultResp_s, *Error_s, error) {
	cResp, err := CreateTask(
		this.clientKey,
		Task_RecaptchaV3TaskProxyless(
			this.websiteUrl,
			this.websiteKey,
			this.minScore,
			this.pageAction,
		),
		0,
		"en",
		"",
	)
	if err != nil {
		return nil, nil, err
	}
	err_s := IsError(cResp)
	if err_s != nil {
		return nil, err_s, nil
	}
	result, err_s, err := Adv_WaitResult(this.clientKey, *cResp)
	if err != nil {
		return nil, nil, err
	}
	if err_s != nil {
		return nil, err_s, nil
	}
	return result, nil, nil
}