package anticaptcha

type GetTaskResultResp_s struct {
	response_s
	Status string `json:"status"`
	Solution Obj `json:"solution"`
	Cost string `json:"cost"`
	Ip string `json:"ip"`
	CreateTime int `json:"createTime"`
	EndTime int `json:"endTime"`
	SolveCount int `json:"solveCount"`
}

func GetTaskResult(clientKey string, taskId int) (*GetTaskResultResp_s, error) {
	resp_s := &GetTaskResultResp_s{}
	err := getApiContent("getTaskResult", Obj{
		"clientKey": clientKey,
		"taskId": taskId,
	}, resp_s)
	if err != nil {
		return nil, err
	}
	return resp_s, nil

	//body, err := json.Marshal(Obj {
	//	"clientKey": clientKey,
	//	"taskId": taskId,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//u := baseURL.ResolveReference(&url.URL{Path: "/getTaskResult"})
	//resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//resp_s := &GetTaskResultResp_s{}
	//err = json.NewDecoder(resp.Body).Decode(resp_s)
	//if err != nil {
	//	return nil, err
	//}
	//return resp_s, nil
}