package anticaptcha

type Adv_SquareNetTextTask struct {
	clientKey string
	body string
	objectName string
	rowsCount int
	columnsCount int
}

func (this *Adv_SquareNetTextTask)New(_clientKey string, _body string, _objectName string, _rowsCount int, _columnsCount int) *Adv_SquareNetTextTask {
	this.clientKey = _clientKey
	this.body = _body
	this.objectName = _objectName
	this.rowsCount = _rowsCount
	this.columnsCount = _columnsCount
	return this
}

func (this *Adv_SquareNetTextTask)Do() (*GetTaskResultResp_s, *Error_s, error) {
	return adv_defaultCreateTask(
		Task_SquareNetTextTask(
			this.body,
			this.objectName,
			this.rowsCount,
			this.columnsCount,
		),
		this.clientKey)
}