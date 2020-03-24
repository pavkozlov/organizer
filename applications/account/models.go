package account

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string     `gorm:"not null;UNIQUE" form:"username" json:"username"`
	Password string     `gorm:"not null" json:"-"`
	Salt     string     `gorm:"not null" json:"-"`
	Session  []Sessions `gorm:"foreignkey:UserID"`
}

type Sessions struct {
	gorm.Model
	UserID       uint      `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	ExpiresIn    time.Time `gorm:"not null"`
	UserAgent    string
}

func (s *Sessions) AfterCreate(tx *gorm.DB) (err error) {
	utc, _ := time.LoadLocation("Europe/Moscow")
	expires := time.Now().Add(time.Hour * 60 * 24).In(utc)
	tx.Model(s).Update("ExpiresIn", expires)
	return
}
