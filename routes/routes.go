package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
	"shopping-system/controllers"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	r.HandleFunc("/home", controllers.Home).Methods("GET")

	buildDir := "./client/build/"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(buildDir, "static")))))
	r.HandleFunc("/{_:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
	})
}
