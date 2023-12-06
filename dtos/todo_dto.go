package dtos

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsFinish    bool   `json:"isFinish"`
}
