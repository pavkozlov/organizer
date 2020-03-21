package account

import (
	"github.com/pavkozlov/organizer/settings"
)

func saveUser(user *User) (err error) {
	if err = settings.Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func getUserByUsername(user *User, username string) (err error) {
	if err = settings.Db.Where("username = ?", username).Find(&user).Error; err != nil {
		return err
	} else {
		return nil
	}
}
