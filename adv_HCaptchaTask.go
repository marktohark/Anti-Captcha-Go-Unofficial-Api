package anticaptcha

type Adv_HCaptchaTask struct {
	clientKey string
	websiteURL string
	websiteKey string
	proxyType ProxyType
	proxyAddress string
	proxyPort int
	proxyLogin string
	proxyPassword string
	userAgent string
}

func (this *Adv_HCaptchaTask)New(_clientKey string, _websiteURL string, _websiteKey string, _proxyType ProxyType, _proxyAddress string, _proxyPort int, _proxyLogin string, _proxyPassword string, _userAgent string) *Adv_HCaptchaTask {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.websiteKey = _websiteKey
	this.proxyType = _proxyType
	this.proxyAddress = _proxyAddress
	this.proxyPort = _proxyPort
	this.proxyLogin = _proxyLogin
	this.proxyPassword = _proxyPassword
	this.userAgent = _userAgent
	return this
}

func (this *Adv_HCaptchaTask)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_HCaptchaTask(
			this.websiteURL,
			this.websiteKey,
			string(this.proxyType),
			this.proxyAddress,
			this.proxyPort,
			this.proxyLogin,
			this.proxyPassword,
			this.userAgent,
			"",
		),
		this.clientKey)
}