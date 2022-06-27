package model

type CalcRequest struct {
	Long  int `json:"long"`
	Width int `json:"width"`
}

type CalcResponse struct {
	Area int `json:"area"`
}
