package account

type User struct {
	ID       uint   `gorm:"AUTO_INCREMENT;UNIQUE;PRIMARY_KEY;" json:"id"`
	Username string `gorm:"UNIQUE" form:"username" json:"username" binding:"required"`
	Password string `gorm:"NOT NULL" json:"-"`
	Salt     string `gorm:"NOT NULL" form:"-" json:"-"`
}
