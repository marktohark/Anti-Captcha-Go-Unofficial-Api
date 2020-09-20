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
	TaskId int
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
	resp_s.TaskId = taskId
	return resp_s, nil
}