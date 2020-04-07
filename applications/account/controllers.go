package account

import (
	"github.com/pavkozlov/organizer/organizer"
)

// SQL для получения юзера по рефреш токену
const refreshTokenQuery = `
	SELECT 
		sessions.deleted_at,
		sessions.id,
		expires_in, 
		users.id AS user_id, 
		users.username
	FROM public.sessions 
	JOIN users ON user_id = users.id
	WHERE refresh_token = ?
	AND sessions.deleted_at IS null
	`

// Получение данных по рефреш токену
func getRefreshToken(rt *refreshTokenRaw, refreshToken string) (err error) {
	if err = organizer.Db.Raw(refreshTokenQuery, refreshToken).Scan(&rt).Error; err != nil {
		return err
	}
	return nil
}

// Отметка токена как удалённого
func deleteRefreshToken(id uint) (err error) {
	if err = organizer.Db.Where("id = ?", id).Delete(Sessions{}).Error; err != nil {
		return err
	}
	return nil
}

// Создание рефреш токена
func createRefreshToken(s *Sessions) (err error) {
	if err = organizer.Db.Create(&s).Error; err != nil {
		return err
	}
	return nil
}

// Создание или получение уже имеющегося рефшер токена
func getOrCrateRefreshToken(s *Sessions, token string) (err error) {
	if err = organizer.Db.Where(&s).Attrs(Sessions{RefreshToken: token}).FirstOrCreate(&s).Error; err != nil {
		return err
	}
	return nil
}

// Создание пользователя
func createUser(user *User) (err error) {
	if err = organizer.Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// Получение пользователя
func getUser(user *User, username string) (err error) {
	if err = organizer.Db.Where("username = ?", username).Find(&user).Error; err != nil {
		return err
	}
	return nil

}
