package controllers

import (
	"Auth/Cms_Authentication_Jwt/api/models"
	"Auth/Cms_Authentication_Jwt/api/utils"
	"net/http"
)

func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	jwtParams, err := utils.JwtExtract(r)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}
	email, ok := jwtParams["user_email"].(string)
	if !ok {
		utils.ToJson(w, "Payload invalid", http.StatusUnauthorized)
		return
	}
	user := models.GetUserByEmail(email)
	if user.Id == 0 {
		utils.ToJson(w, "User not found", http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, user, http.StatusOK)
}
