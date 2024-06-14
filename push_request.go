package models

type PushRequest struct {
	Id     string            `json:"id"`
	Body   string            `json:"body"`
	Title  string            `json:"title"`
	Tokens []string          `json:"tokens"`
	Data   map[string]string `json:"data"`
}

type PushResponse struct {
	Id        string `json:"id"`
	Token     string `json:"token"`
	Success   bool   `json:"success"`
	MessageId string `json:"messageId"`
	Error     string `json:"error"`
}
