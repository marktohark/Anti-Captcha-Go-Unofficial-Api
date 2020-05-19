package anticaptcha

type Adv_FunCaptchaTaskProxyless struct {
	clientKey string
	websiteURL string
	funcaptchaApiJSSubdomain string
	data string
	websitePublicKey string
}

func (this *Adv_FunCaptchaTaskProxyless)New(_clientKey string, _websiteURL string, _funcaptchaApiJSSubdomain string, _data string, _websitePublicKey string) *Adv_FunCaptchaTaskProxyless {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.funcaptchaApiJSSubdomain = _funcaptchaApiJSSubdomain
	this.data = _data
	this.websitePublicKey = _websitePublicKey
	return this
}

func (this *Adv_FunCaptchaTaskProxyless)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_FunCaptchaTaskProxyless(
			this.websiteURL,
			this.funcaptchaApiJSSubdomain,
			this.data,
			this.websitePublicKey,
		),
		this.clientKey)
}