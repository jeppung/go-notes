package models

type Note struct {
	BaseModel
	Title string `json:"title"`
	Body  string `json:"body"`
}
