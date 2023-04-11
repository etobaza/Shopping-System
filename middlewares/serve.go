package middlewares

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	buildDir := "./client/build/"
	file := filepath.Join(buildDir, "index.html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(w, r, file)
}

func SetupStaticFileServer(r *mux.Router) {
	buildDir := "./client/build/"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(buildDir, "static")))))
	r.HandleFunc("/{_:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
	})
}

func ServeMessage() {
	fmt.Println("   ┌───────────────────────────────────────────┐\n   │                                           │\n   │   Serving!                                │\n   │                                           │\n   │   - Golang: http://localhost:8080         │\n   │   - React:  http://localhost:3000         │\n   │                                           │\n   │   Don't forget to launch React app!       │\n   │                                           │\n   └───────────────────────────────────────────┘")
}
