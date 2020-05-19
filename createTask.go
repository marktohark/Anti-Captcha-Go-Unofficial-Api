package anticaptcha

type CreateTaskResp_s struct {
	response_s
	TaskId int `json:"taskId"`
}

func CreateTask(clientKey string, task Obj, softId int, languagePool string, callbackUrl string) (*CreateTaskResp_s, error) {
	resp_s := &CreateTaskResp_s{}
	err := getApiContent("createTask", Obj{
		"clientKey": clientKey,
		"task": task,
		"softId": softId,
		"languagePool": languagePool,
		"callbackUrl": callbackUrl,
	}, resp_s)
	if err != nil {
		return nil, err
	}
	return resp_s, nil

	//body, err := json.Marshal(Obj {
	//	"clientKey": clientKey,
	//	"task": task,
	//	"softId": softId,
	//	"languagePool": languagePool,
	//	"callbackUrl": callbackUrl,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//u := baseURL.ResolveReference(&url.URL{Path: "/createTask"})
	//resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//resp_s := &CreateTaskResp_s{}
	//err = json.NewDecoder(resp.Body).Decode(resp_s)
	//if err != nil {
	//	return nil, err
	//}
	//return resp_s, nil
}

func Task_ImageToTextTask(body string, phrase bool, _case bool, numeric int, math bool, minLength int, maxLength int, comment string, websiteURL string) Obj {
	return Obj {
		"type": "ImageToTextTask",
		"body":body,
		"phrase":phrase,
		"case":_case,
		"numeric":numeric,
		"math":math,
		"minLength":minLength,
		"maxLength":maxLength,
		"comment": comment,
		"websiteURL": websiteURL,
	}
}

func Task_NoCaptchaTask(websiteURL string, websiteKey string, websiteSToken string, proxyType string, proxyAddress string, proxyPort int, proxyLogin string, proxyPassword string, userAgent string, cookies string, isInvisible bool) Obj {
	return Obj {
		"type": "NoCaptchaTask",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
		"websiteSToken": websiteSToken,
		"proxyType": proxyType,
		"proxyAddress": proxyAddress,
		"proxyPort": proxyPort,
		"proxyLogin": proxyLogin,
		"proxyPassword": proxyPassword,
		"userAgent": userAgent,
		"cookies": cookies,
		"isInvisible": isInvisible,
	}
}

func Task_NoCaptchaTaskProxyless(websiteURL string, websiteKey string, websiteSToken string, isInvisible bool) Obj {
	return Obj {
		"type": "NoCaptchaTaskProxyless",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
		"websiteSToken": websiteSToken,
		"isInvisible": isInvisible,
	}
}

func Task_RecaptchaV3TaskProxyless(websiteURL string, websiteKey string, minScore float64, pageAction string) Obj {
	return Obj {
		"type": "RecaptchaV3TaskProxyless",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
		"minScore": minScore,
		"pageAction": pageAction,
	}
}

func Task_FunCaptchaTask(websiteURL string, funcaptchaApiJSSubdomain string, data string, websitePublicKey string, proxyType string, proxyAddress string, proxyPort int, proxyLogin string, proxyPassword string, userAgent string, cookies string) Obj {
	return Obj {
		"type": "FunCaptchaTask",
		"websiteURL": websiteURL,
		"funcaptchaApiJSSubdomain": funcaptchaApiJSSubdomain,
		"data": data,
		"websitePublicKey": websitePublicKey,
		"proxyType": proxyType,
		"proxyAddress": proxyAddress,
		"proxyPort": proxyPort,
		"proxyLogin": proxyLogin,
		"proxyPassword": proxyPassword,
		"userAgent": userAgent,
		"cookies": cookies,
	}
}

func Task_FunCaptchaTaskProxyless(websiteURL string, funcaptchaApiJSSubdomain string, data string, websitePublicKey string) Obj {
	return Obj {
		"type": "FunCaptchaTaskProxyless",
		"websiteURL": websiteURL,
		"funcaptchaApiJSSubdomain": funcaptchaApiJSSubdomain,
		"data": data,
		"websitePublicKey": websitePublicKey,
	}
}

func Task_SquareNetTextTask(body string, objectName string, rowsCount int, columnsCount int) Obj {
	return Obj {
		"type": "SquareNetTextTask",
		"body": body,
		"objectName": objectName,
		"rowsCount": rowsCount,
		"columnsCount": columnsCount,
	}
}

func Task_GeeTestTask(websiteURL string, gt string, challenge string, geetestApiServerSubdomain string, proxyType string, proxyAddress string, proxyPort int, proxyLogin string, proxyPassword string, userAgent string, cookies string) Obj {
	return Obj {
		"type": "GeeTestTask",
		"websiteURL": websiteURL,
		"gt": gt,
		"challenge": challenge,
		"geetestApiServerSubdomain": geetestApiServerSubdomain,
		"proxyType": proxyType,
		"proxyAddress": proxyAddress,
		"proxyPort": proxyPort,
		"proxyLogin": proxyLogin,
		"proxyPassword": proxyPassword,
		"userAgent": userAgent,
		"cookies": cookies,
	}
}

func Task_GeeTestTaskProxyless(websiteURL string, gt string, challenge string, geetestApiServerSubdomain string) Obj {
	return Obj {
		"type": "GeeTestTaskProxyless",
		"websiteURL": websiteURL,
		"gt": gt,
		"challenge": challenge,
		"geetestApiServerSubdomain": geetestApiServerSubdomain,
	}
}

func Task_HCaptchaTask(websiteURL string, websiteKey string, proxyType string, proxyAddress string, proxyPort int, proxyLogin string, proxyPassword string, userAgent string, cookies string) Obj {
	return Obj {
		"type": "HCaptchaTask",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
		"proxyType": proxyType,
		"proxyAddress": proxyAddress,
		"proxyPort": proxyPort,
		"proxyLogin": proxyLogin,
		"proxyPassword": proxyPassword,
		"userAgent": userAgent,
		"cookies": cookies,
	}
}

func Task_HCaptchaTaskProxyless(websiteURL string, websiteKey string) Obj {
	return Obj {
		"type": "HCaptchaTaskProxyless",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
}

func Task_RecaptchaV1Task(websiteURL string, websiteKey string, proxyType string, proxyAddress string, proxyPort int, proxyLogin string, proxyPassword string, userAgent string, cookies string) Obj {
	return Obj {
		"type": "RecaptchaV1Task",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
		"proxyType": proxyType,
		"proxyAddress": proxyAddress,
		"proxyPort": proxyPort,
		"proxyLogin": proxyLogin,
		"proxyPassword": proxyPassword,
		"userAgent": userAgent,
		"cookies": cookies,
	}
}

func Task_RecaptchaV1TaskProxyless(websiteURL string, websiteKey string) Obj {
	return Obj {
		"type": "RecaptchaV1TaskProxyless",
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
}
