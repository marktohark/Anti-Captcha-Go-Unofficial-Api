package anticaptcha

type response_i interface {
	GetErrorId() int
	SetErrorCode(string)
	SetErrorDesc(string)
}

type response_s struct {
	response_i
	ErrorId int `json:"errorId"`
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDescription"`
}

func (this *response_s) GetErrorId() int {
	return this.ErrorId
}

func (this *response_s) SetErrorCode(code string) {
	this.ErrorCode = code
}

func (this *response_s) SetErrorDesc(desc string) {
	this.ErrorDesc = desc
}

type Error_s response_s

func IsError(resp response_i) *Error_s {
	errId := resp.GetErrorId()
	if errId == 0  {
		return nil
	}
	temp, has := errorMap[errId]
	if !has {
		panic("can't find errorId in the errorMap")
	}
	return &temp
}

var errorMap map[int]Error_s = map[int]Error_s {
	1: {ErrorId: 1, ErrorCode: "ERROR_KEY_DOES_NOT_EXIST", ErrorDesc: `Account authorization key not found in the system`},
	2: {ErrorId: 2, ErrorCode: "ERROR_NO_SLOT_AVAILABLE", ErrorDesc: `No idle captcha workers are available at the moment, please try a bit later or try increasing your maximum bid here`},
	3: {ErrorId: 3, ErrorCode: "ERROR_ZERO_CAPTCHA_FILESIZE", ErrorDesc: `The size of the captcha you are uploading is less than 100 bytes.`},
	4: {ErrorId: 4, ErrorCode: "ERROR_TOO_BIG_CAPTCHA_FILESIZE", ErrorDesc: `The size of the captcha you are uploading is more than 500,000 bytes.`},
	10: {ErrorId: 10, ErrorCode: "ERROR_ZERO_BALANCE", ErrorDesc: `Account has zeo or negative balance`},
	11: {ErrorId: 11, ErrorCode: "ERROR_IP_NOT_ALLOWED", ErrorDesc: `Request with current account key is not allowed from your IP. Please refer to IP list section located here`},
	12: {ErrorId: 12, ErrorCode: "ERROR_CAPTCHA_UNSOLVABLE", ErrorDesc: `Captcha could not be solved by 5 different workers`},
	13: {ErrorId: 13, ErrorCode: "ERROR_BAD_DUPLICATES", ErrorDesc: `100% recognition feature did not work due to lack of amount of guess attempts`},
	14: {ErrorId: 14, ErrorCode: "ERROR_NO_SUCH_METHOD", ErrorDesc: `Request to API made with method which does not exist`},
	15: {ErrorId: 15, ErrorCode: "ERROR_IMAGE_TYPE_NOT_SUPPORTED", ErrorDesc: `Could not determine captcha file type by its exif header or image type is not supported. The only allowed formats are JPG, GIF, PNG`},
	16: {ErrorId: 16, ErrorCode: "ERROR_NO_SUCH_CAPCHA_ID", ErrorDesc: `Captcha you are requesting does not exist in your current captchas list or has been expired. Captchas are removed from API after 5 minutes after upload.`},
	20: {ErrorId: 20, ErrorCode: "ERROR_EMPTY_COMMENT", ErrorDesc: `"comment" property is required for this request`},
	21: {ErrorId: 21, ErrorCode: "ERROR_IP_BLOCKED", ErrorDesc: `Your IP is blocked due to API inproper use. Check the reason at https://anti-captcha.com/panel/tools/ipsearch`},
	22: {ErrorId: 22, ErrorCode: "ERROR_TASK_ABSENT", ErrorDesc: `Task property is empty or not set in createTask method. Please refer to API v2 documentation.`},
	23: {ErrorId: 23, ErrorCode: "ERROR_TASK_NOT_SUPPORTED", ErrorDesc: `Task type is not supported or inproperly printed. Please check \"type\" parameter in task object.`},
	24: {ErrorId: 24, ErrorCode: "ERROR_INCORRECT_SESSION_DATA", ErrorDesc: `Some of the required values for successive user emulation are missing.`},
	25: {ErrorId: 25, ErrorCode: "ERROR_PROXY_CONNECT_REFUSED", ErrorDesc: `Could not connect to proxy related to the task, connection refused`},
	26: {ErrorId: 26, ErrorCode: "ERROR_PROXY_CONNECT_TIMEOUT", ErrorDesc: `Could not connect to proxy related to the task, connection timeout`},
	27: {ErrorId: 27, ErrorCode: "ERROR_PROXY_READ_TIMEOUT", ErrorDesc: `Connection to proxy for task has timed out`},
	28: {ErrorId: 28, ErrorCode: "ERROR_PROXY_BANNED", ErrorDesc: `Proxy IP is banned by target service`},
	29: {ErrorId: 29, ErrorCode: "ERROR_PROXY_TRANSPARENT", ErrorDesc: `Task denied at proxy checking state. Proxy must be non-transparent to hide our server IP.`},
	30: {ErrorId: 30, ErrorCode: "ERROR_RECAPTCHA_TIMEOUT", ErrorDesc: `Recaptcha task timeout, probably due to slow proxy server or Google server`},
	31: {ErrorId: 31, ErrorCode: "ERROR_RECAPTCHA_INVALID_SITEKEY", ErrorDesc: `Recaptcha server reported that site key is invalid`},
	32: {ErrorId: 32, ErrorCode: "ERROR_RECAPTCHA_INVALID_DOMAIN", ErrorDesc: `Recaptcha server reported that domain for this site key is invalid`},
	33: {ErrorId: 33, ErrorCode: "ERROR_RECAPTCHA_OLD_BROWSER", ErrorDesc: `Recaptcha server reported that browser user-agent is not compatible with their javascript`},
	34: {ErrorId: 34, ErrorCode: "ERROR_TOKEN_EXPIRED", ErrorDesc: `Captcha provider server reported that additional variable token has been expired. Please try again with new token.`},
	35: {ErrorId: 35, ErrorCode: "ERROR_PROXY_HAS_NO_IMAGE_SUPPORT", ErrorDesc: `Proxy does not support transfer of image data from Google servers`},
	36: {ErrorId: 36, ErrorCode: "ERROR_PROXY_INCOMPATIBLE_HTTP_VERSION", ErrorDesc: `Proxy does not support long GET requests with length about 2000 bytes and does not support SSL connections`},
	37: {ErrorId: 37, ErrorCode: "ERROR_FACTORY_SERVER_API_CONNECTION_FAILED", ErrorDesc: `Could not connect to Factory Server API within 5 seconds`},
	38: {ErrorId: 38, ErrorCode: "ERROR_FACTORY_SERVER_BAD_JSON", ErrorDesc: `Incorrect Factory Server JSON response, something is broken`},
	39: {ErrorId: 39, ErrorCode: "ERROR_FACTORY_SERVER_ERRORID_MISSING", ErrorDesc: `Factory Server API did not send any errorId`},
	40: {ErrorId: 40, ErrorCode: "ERROR_FACTORY_SERVER_ERRORID_NOT_ZERO", ErrorDesc: `Factory Server API reported errorId != 0, check this error`},
	41: {ErrorId: 41, ErrorCode: "ERROR_FACTORY_MISSING_PROPERTY", ErrorDesc: `Some of the required property values are missing in Factory form specifications. Customer must send all required values.`},
	42: {ErrorId: 42, ErrorCode: "ERROR_FACTORY_PROPERTY_INCORRECT_FORMAT", ErrorDesc: `Expected other type of property value in Factory form structure. Customer must send specified value type.`},
	43: {ErrorId: 43, ErrorCode: "ERROR_FACTORY_ACCESS_DENIED", ErrorDesc: `Factory control belong to another account, check your account key.`},
	44: {ErrorId: 44, ErrorCode: "ERROR_FACTORY_SERVER_OPERATION_FAILED", ErrorDesc: `Factory Server general error code`},
	45: {ErrorId: 45, ErrorCode: "ERROR_FACTORY_PLATFORM_OPERATION_FAILED", ErrorDesc: `Factory Platform general error code.`},
	46: {ErrorId: 46, ErrorCode: "ERROR_FACTORY_PROTOCOL_BROKEN", ErrorDesc: `Factory task lifetime protocol broken during task workflow.`},
	47: {ErrorId: 47, ErrorCode: "ERROR_FACTORY_TASK_NOT_FOUND", ErrorDesc: `Task not found or not available for this operation`},
	48: {ErrorId: 48, ErrorCode: "ERROR_FACTORY_IS_SANDBOXED", ErrorDesc: `Factory is sandboxed, creating tasks is possible only by Factory owner. Switch it to production mode to make it available for other customers.`},
	49: {ErrorId: 49, ErrorCode: "ERROR_PROXY_NOT_AUTHORISED", ErrorDesc: `Proxy login and password are incorrect`},
	50: {ErrorId: 50, ErrorCode: "ERROR_FUNCAPTCHA_NOT_ALLOWED", ErrorDesc: `Customer did not enable Funcaptcha Proxyless tasks in Customers Area - API Settings. All customers must read terms, pass mini test and sign/accept the form before being able to use this feature.`},
	51: {ErrorId: 51, ErrorCode: "ERROR_INVISIBLE_RECAPTCHA", ErrorDesc: `Recaptcha was attempted to be solved as usual one, instead of invisible mode. Basically you don't need to do anything when this error occurs, just continue sending tasks with this domain. Our system will self-learn to solve recaptchas from this sitekey in invisible mode.`},
	52: {ErrorId: 52, ErrorCode: "ERROR_FAILED_LOADING_WIDGET", ErrorDesc: `Could not load captcha provider widget in worker browser. Please try sending new task.`},
}