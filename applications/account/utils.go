package account

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/pavkozlov/organizer/settings"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateSalt(saltLen int) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	salt := ""
	for i := 0; i < saltLen; i++ {
		salt += string(charset[seededRand.Intn(len(charset))])
	}
	return salt
}

func encryptPassword(password, salt string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(password + salt))
	encryptedPassword := sha_512.Sum([]byte(""))
	return hex.EncodeToString(encryptedPassword)
}

func authorize(username, password string) bool {
	user := User{}
	e := settings.Db.Where("username = ?", username).Find(&user)
	if e.Error != nil {
		return false
	}

	if user.Password == encryptPassword(password, user.Salt) {
		return true
	} else {
		return false
	}

}
