package account

import "time"

// Форма для регистрации / логина
type userForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Форма для обновления токена
type refreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Форма для ответа на получение данных по рефреш токену
type refreshTokenRaw struct {
	ID, UserID uint
	ExpiresIn  time.Time
	Username   string
}
