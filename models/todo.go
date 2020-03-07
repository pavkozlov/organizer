package models

import "time"

type Todo struct {
	ID        uint      `gorm:"AUTO_INCREMENT;UNIQUE;PRIMARY_KEY;" json:"id"`
	Title     string    `gorm:"NOT NULL" form:"title" json:"title" binding:"required"`
	Complete  bool      `gorm:"DEFAULT:false" json:"complete"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created"`
}
