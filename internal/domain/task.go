package domain

type Task struct {
	ID       int `json:"id"`
	Text     string `json:"text"`
	Datetime string `json:"datetime"`
	Status   string `json:"status"`
}
