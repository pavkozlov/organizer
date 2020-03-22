package account

import (
	"github.com/pavkozlov/organizer/settings"
	"time"
)

const REFRESHTOKEN_QUERY = `
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

type refreshTokenRaw struct {
	ID, UserID uint
	ExpiresIn  time.Time
	Username   string
}

func getRefreshToken(rt *refreshTokenRaw, refreshToken string) (err error) {
	if err = settings.Db.Raw(REFRESHTOKEN_QUERY, refreshToken).Scan(&rt).Error; err != nil {
		return err
	}
	return nil
}

func deleteRefreshToken(id uint) (err error) {
	if err = settings.Db.Where("id = ?", id).Delete(Sessions{}).Error; err != nil {
		return err
	}
	return nil
}

func createRefreshToken(s *Sessions) (err error) {
	if err = settings.Db.Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func getOrCrateRefreshToken(s *Sessions, token string) (err error) {
	if err = settings.Db.Where(&s).Attrs(Sessions{RefreshToken: token}).FirstOrCreate(&s).Error; err != nil {
		return err
	}
	return nil
}

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
