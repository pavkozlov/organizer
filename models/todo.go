package models

type Todo struct {
	ID    uint   `json:"id"`
	Title string `form:"title" json:"title" binding:"required"`
}
