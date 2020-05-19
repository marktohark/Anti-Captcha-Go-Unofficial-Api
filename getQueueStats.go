package anticaptcha

type GetQueueStatsResp_s struct {
	Waiting int `json:"waiting"`
	Load float32 `json:"load"`
	Bid float32 `json:"bid"`
	Speed float32 `json:"speed"`
	Total int `json:"total"`
}

const (
	QueueId_ImageToText_Eng = 1
	QueueId_ImageToText_Russian = 2
	QueueId_NoCaptcha = 5
	QueueId_NoCaptcha_Proxyless = 6
	QueueId_Funcaptcha = 7
	QueueId_Funcaptcha_Proxyless = 10
	QueueId_SquareNetText = 11
	QueueId_GeeTest = 12
	QueueId_GeeTest_Proxyless = 13
	QueueId_RecaptchaV3_3 = 18
	QueueId_RecaptchaV3_7 = 19
	QueueId_RecaptchaV3_9 = 20
	QueueId_hCaptcha = 21
	QueueId_hCaptcha_Proxyless = 22
)

type QueueId int

func getQueueStats(queueId QueueId) (*GetQueueStatsResp_s, error) {
	resp_s := &GetQueueStatsResp_s{}
	err := getApiContent("getQueueStats", Obj{
		"queueId": queueId,
	}, resp_s)
	if err != nil {
		return nil, err
	}
	return resp_s, nil
}