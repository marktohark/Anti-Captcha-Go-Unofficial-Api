package anticaptcha

type Adv_NoCaptchaTaskProxyless struct {
	clientKey string
	websiteURL string
	websiteKey string
}

func (this *Adv_NoCaptchaTaskProxyless)New(_clientKey string, _websiteURL string, _websiteKey string) *Adv_NoCaptchaTaskProxyless {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.websiteKey = _websiteKey
	return this
}

func (this *Adv_NoCaptchaTaskProxyless)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_NoCaptchaTaskProxyless(
			this.websiteURL,
			this.websiteKey,
			"",
			false,
		),
		this.clientKey)
}