package auth

import (
	"Auth/Cms_Authentication_Jwt/api/models"
	"Auth/Cms_Authentication_Jwt/api/security"
	"Auth/Cms_Authentication_Jwt/api/utils"
	"errors"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

func SignIn(email, password string) (string, error) {
	user := models.GetUserByEmail(email)
	if user.Id == 0 {
		return "", ErrUserNotFound
	}
	err := security.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
