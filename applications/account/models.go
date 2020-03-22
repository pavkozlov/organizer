package account

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"not null;UNIQUE" form:"username" json:"username"`
	Password string  `gorm:"not null" json:"-"`
	Salt     string  `gorm:"not null" json:"-"`
	Tokens   []Token `gorm:"foreignkey:UserID;association_foreignkey:TokenID"`
}

type Token struct {
	gorm.Model
	UserID       uint
	RefreshToken string
	ExpiresIn    int
}
