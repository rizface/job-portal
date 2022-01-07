package response

type Standard struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data map[string]interface{} `json:"data"`
}
