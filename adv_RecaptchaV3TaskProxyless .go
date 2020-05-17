package anticaptcha

type Adv_RecaptchaV3TaskProxyless struct {
	clientKey string
	websiteUrl string
	websiteKey string
	minScore float64
	pageAction string
}

func (this *Adv_RecaptchaV3TaskProxyless)New(_clientKey string, _websiteUrl string, _websiteKey string, _minScore float64, pageAction string) *Adv_RecaptchaV3TaskProxyless {
	this.clientKey = _clientKey
	this.websiteUrl = _websiteUrl
	this.websiteKey = _websiteKey
	this.minScore = _minScore
	this.pageAction = pageAction
	return this
}

func (this *Adv_RecaptchaV3TaskProxyless)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_RecaptchaV3TaskProxyless(
			this.websiteUrl,
			this.websiteKey,
			this.minScore,
			this.pageAction,
		),
		this.clientKey)
}