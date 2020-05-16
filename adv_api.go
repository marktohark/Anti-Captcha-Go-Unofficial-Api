package anticaptcha

import (
	"fmt"
	"time"
)

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