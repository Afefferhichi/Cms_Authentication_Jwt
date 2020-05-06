package controllers

import (
	"Auth/Cms_Authentication_Jwt/api/models"
	"Auth/Cms_Authentication_Jwt/api/utils"
	"encoding/json"
	"net/http"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	rs, err := models.CreateUser(user)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rs, http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	utils.ToJson(w, users, http.StatusCreated)
}
