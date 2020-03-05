package models

import "time"

type Todo struct {
	ID    uint   `json:"id"`
	Title string `form:"title" json:"title" binding:"required"`
	Complete bool `default:"false" json:"complete"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created"`
}
