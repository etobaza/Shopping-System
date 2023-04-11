package routes

import (
	"github.com/gorilla/mux"
	"shopping-system/controllers"
	"shopping-system/middlewares"
)

func SetupRoutes(r *mux.Router) {

	// All users routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	r.HandleFunc("/shop", controllers.Shop).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	// Admin routes
	r.HandleFunc("/admin-panel", controllers.AdminPanel).Methods("GET")

	// Misc routes
	r.HandleFunc("/items", controllers.GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	r.HandleFunc("/shop/items", controllers.GetItemsSeller).Methods("GET")
	r.HandleFunc("/shop/items", controllers.CreateItem).Methods("POST")
	r.HandleFunc("/shop/items/{id}", controllers.DeleteItem).Methods("DELETE")

	middlewares.SetupStaticFileServer(r)
}
