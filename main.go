package main

import (
	"goapp/auth"
	"goapp/users"
	"goapp/views"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	logger := slog.Default()
	logger.Info("Starting server")

	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(views.Home()))

	fileHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	mux.Handle("/assets/", fileHandler)

	mux.Handle("/api/v1/user/", auth.Authenticate(users.UserMux()))

	http.ListenAndServe(":8080", mux)
	logger.Info("Server stopped")
}
