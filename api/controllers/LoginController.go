package controllers

import (
	"Auth/Cms_Authentication_Jwt/api/auth"
	"Auth/Cms_Authentication_Jwt/api/models"
	"Auth/Cms_Authentication_Jwt/api/utils"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, token, http.StatusOK)
}
