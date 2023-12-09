package web

type CategoryResponse struct {//isinya cuma id dan name karena mengikuti apispec.json untuk response nya
	Id   int    `json:"id"`
	Name string	`json:"name"`
}
