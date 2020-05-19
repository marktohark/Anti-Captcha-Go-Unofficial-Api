package anticaptcha

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Obj map[string]interface{}

var (
	baseURL       = &url.URL{Host: "api.anti-captcha.com", Scheme: "https", Path: "/"}
)

func getApiContent(apiName string, obj Obj, result interface{}) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	u := baseURL.ResolveReference(&url.URL{Path: "/" + apiName})
	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}