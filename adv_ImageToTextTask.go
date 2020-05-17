package anticaptcha

const (
	Adv_imageNumeric_NoRequire = 0 //no requirements
	Adv_imageNumeric_OnlyNumber = 1 //only number are allowed
	Adv_imageNumeric_ExceptNumber = 2 //any letters are allowed except numbers
)

type ImageNumeric int

type Adv_ImageToTextTask struct {
	clientKey string
	body string
	_case bool
	numeric ImageNumeric
}

func (this *Adv_ImageToTextTask)New(_clientKey string, _body string, __case bool, _numeric ImageNumeric) *Adv_ImageToTextTask {
	this.clientKey = _clientKey
	this.body = _body
	this._case = __case
	this.numeric = _numeric
	return this
}

func (this *Adv_ImageToTextTask)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_ImageToTextTask(
			this.body,
			false,
			this._case,
			int(this.numeric),
			false,
			0,
			0,
			"",
			"",
		),
		this.clientKey)
}