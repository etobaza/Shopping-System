package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"shopping-system/config"
	"shopping-system/middlewares"
	"shopping-system/routes"
)

func main() {
	config.InitDB()
	r := mux.NewRouter()
	r.Use(middlewares.CorsMiddleware)
	routes.SetupRoutes(r)
	middlewares.ServeMessage()
	http.ListenAndServe(":8080", r)
}
