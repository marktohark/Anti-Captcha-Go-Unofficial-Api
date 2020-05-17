package anticaptcha

type Adv_FunCaptchaTask struct {
	clientKey string
	websiteURL string
	funcaptchaApiJSSubdomain string
	data string
	websitePublicKey string
	proxyType ProxyType
	proxyAddress string
	proxyPort int
	proxyLogin string
	proxyPassword string
	userAgent string
}

func (this *Adv_FunCaptchaTask)New(_clientKey string, _websiteURL string, _funcaptchaApiJSSubdomain string, _data string, _websitePublicKey string, _proxyType ProxyType, _proxyAddress string, _proxyPort int, _proxyLogin string, _proxyPassword string, _userAgent string) *Adv_FunCaptchaTask {
	this.clientKey = _clientKey
	this.websiteURL = _websiteURL
	this.funcaptchaApiJSSubdomain = _funcaptchaApiJSSubdomain
	this.data = _data
	this.websitePublicKey = _websitePublicKey
	this.proxyType = _proxyType
	this.proxyAddress = _proxyAddress
	this.proxyPort = _proxyPort
	this.proxyLogin = _proxyLogin
	this.proxyPassword = _proxyPassword
	this.userAgent = _userAgent
	return this
}

func (this *Adv_FunCaptchaTask)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_FunCaptchaTask(
			this.websiteURL,
			this.funcaptchaApiJSSubdomain,
			this.data,
			this.websitePublicKey,
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