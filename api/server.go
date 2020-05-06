package api

import (
	"Auth/Cms_Authentication_Jwt/api/models"
	"Auth/Cms_Authentication_Jwt/api/routes"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	db := models.Connect()
	if !db.HasTable(&models.User{}) {
		db.Debug().CreateTable(&models.User{})
	}
	db.Close()
	listen(8000)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("Listening Port %s...\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, routes.LoadCors(r)))
}
