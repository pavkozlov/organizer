package account

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pavkozlov/organizer/organizer"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Генерация рандомной строки заданной длины
func generateRandomString(strtLen int) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	salt := ""
	for i := 0; i < strtLen; i++ {
		salt += string(charset[seededRand.Intn(len(charset))])
	}
	return salt
}

// Зашифровать пароль
func encryptPassword(password, salt string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(password + salt))
	encryptedPassword := sha_512.Sum([]byte(""))
	return hex.EncodeToString(encryptedPassword)
}

// Авторизация. Находит пользователя в БД, сверяет пароль
func authorize(username, password string, user *User) bool {

	if getUser(user, username) != nil {
		return false
	}

	if user.Password == encryptPassword(password, user.Salt) {
		return true
	} else {
		return false
	}

}

// Генерация аццесс токена
func generateToken(username string, id uint) (t string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"expired":  time.Now().Add(time.Minute * 30).Unix(),
	})
	t, _ = token.SignedString([]byte(organizer.SecretKey))
	return t
}
