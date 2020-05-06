package routes

import (
	"Auth/Cms_Authentication_Jwt/api/controllers"
	"Auth/Cms_Authentication_Jwt/api/middlewares"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/profile", middlewares.IsAuth(controllers.ProtectedRoute)).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/register", controllers.PostUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	return r
}
