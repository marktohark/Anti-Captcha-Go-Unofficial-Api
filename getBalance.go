package anticaptcha

type GetBalanceResp_s struct {
	response_s
	Balance int `json:"balance"`
}

func GetBalance(clientKey string) (*GetBalanceResp_s, error) {
	resp_s := &GetBalanceResp_s{}
	err := getApiContent("getBalance", Obj{
		"clientKey": clientKey,
	}, resp_s)
	if err != nil {
		return nil, err
	}
	return resp_s, nil
}