package web

type WebResponse struct {
	Code int	`json:"code"`
	Status string `json:"string"`
	Data interface{} `json:"data"`
}