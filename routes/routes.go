package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"shopping-system/controllers"
)

func SetupRoutes(r *mux.Router, uc *controllers.UserController) {
	fileServer := http.FileServer(http.Dir("./views"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	r.HandleFunc("/register", uc.GetRegister).Methods("GET")
	r.HandleFunc("/register", uc.Register).Methods("POST")
	r.HandleFunc("/login", uc.GetLogin).Methods("GET")
	r.HandleFunc("/login", uc.Login).Methods("POST")
}
