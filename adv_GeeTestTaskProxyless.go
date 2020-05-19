package anticaptcha

type Adv_GeeTestTaskProxyless struct {
	clientKey string
	websiteURL string
	gt string
	challenge string
}

func (this *Adv_GeeTestTaskProxyless)New(_clientKey string, _websiteURL string, _gt string, _challenge string) *Adv_GeeTestTaskProxyless {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.gt = _gt
	this.challenge = _challenge
	return this
}

func (this *Adv_GeeTestTaskProxyless)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_GeeTestTaskProxyless(
			this.websiteURL,
			this.gt,
			this.challenge,
			"",
		),
		this.clientKey)
}