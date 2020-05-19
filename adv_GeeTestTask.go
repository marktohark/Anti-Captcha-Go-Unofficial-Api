package anticaptcha

type Adv_GeeTestTask struct {
	clientKey string
	websiteURL string
	gt string
	challenge string
	proxyType ProxyType
	proxyAddress string
	proxyPort int
	proxyLogin string
	proxyPassword string
	userAgent string
}

func (this *Adv_GeeTestTask)New(_clientKey string, _websiteURL string, _gt string, _challenge string, _proxyType ProxyType, _proxyAddress string, _proxyPort int, _proxyLogin string, _proxyPassword string, _userAgent string) *Adv_GeeTestTask {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.gt = _gt
	this.challenge = _challenge
	this.proxyType = _proxyType
	this.proxyAddress = _proxyAddress
	this.proxyPort = _proxyPort
	this.proxyLogin = _proxyLogin
	this.proxyPassword = _proxyPassword
	this.userAgent = _userAgent
	return this
}

func (this *Adv_GeeTestTask)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_GeeTestTask(
			this.websiteURL,
			this.gt,
			this.challenge,
			"",
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