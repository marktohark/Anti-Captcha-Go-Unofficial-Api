package anticaptcha

import (
	"net/url"
)

type Obj map[string]interface{}

var (
	baseURL       = &url.URL{Host: "api.anti-captcha.com", Scheme: "https", Path: "/"}
)