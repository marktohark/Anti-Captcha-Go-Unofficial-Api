package anticaptcha

type Adv_HCaptchaTaskProxyless struct {
	clientKey string
	websiteURL string
	websiteKey string
}

func (this *Adv_HCaptchaTaskProxyless)New(_clientKey string, _websiteURL string, _websiteKey string) *Adv_HCaptchaTaskProxyless {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.websiteKey = _websiteKey
	return this
}

func (this *Adv_HCaptchaTaskProxyless)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_HCaptchaTaskProxyless(
			this.websiteURL,
			this.websiteKey,
		),
		this.clientKey)
}