package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
	"shopping-system/controllers"
)

func SetupRoutes(r *mux.Router, uc *controllers.UserController) {
	r.HandleFunc("/register", uc.Register).Methods("POST")
	r.HandleFunc("/login", uc.Login).Methods("POST")

	buildDir := "./client/build/"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(buildDir, "static")))))
	r.HandleFunc("/{_:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
	})
}
