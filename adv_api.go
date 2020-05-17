package anticaptcha

import (
	"fmt"
	"time"
)

const (
	Adv_proxyType_Http = "http"
	Adv_proxyType_Https = "https"
	Adv_proxyType_Socks4 = "socks4"
	Adv_proxyType_Socks5 = "socks5"
)

type ProxyType string

func Adv_WaitResult(clientKey string, cResp CreateTaskResp_s) (*GetTaskResultResp_s, *Error_s, error) {
	check := time.NewTimer(2 * time.Second)
	defer check.Stop()
	for {
		select {
		case <-check.C:
			rResp, err := GetTaskResult(clientKey, cResp.TaskId)
			if err != nil {
				return nil, nil, err
			}
			if err_s := IsError(rResp); err_s != nil {
				return nil, err_s, nil
			}
			switch(rResp.Status) {
			case "processing":
				fmt.Println("Status =>", rResp.Status)
				check.Reset(2 * time.Second)
			case "ready":
				return rResp, nil, nil
			default:
				panic("anti-captcha get new Status, not in (processing and ready)")
			}
		}
	}
}

func adv_defaultCreateTask(Task Obj, clientKey string) (*GetTaskResultResp_s, *Error_s, error) {
	cResp, err := CreateTask(
		clientKey,
		Task,
		0,
		"en",
		"",
	)
	if err != nil {
		return nil, nil, err
	}
	err_s := IsError(cResp)
	if err_s != nil {
		return nil, err_s, nil
	}
	result, err_s, err := Adv_WaitResult(clientKey, *cResp)
	if err != nil {
		return nil, nil, err
	}
	if err_s != nil {
		return nil, err_s, nil
	}
	return result, nil, nil
}