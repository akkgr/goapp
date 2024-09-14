package main

import (
	"embed"
	"goapp/auth"
	"goapp/users"
	"goapp/views"
	"io/fs"
	"log/slog"
	"net/http"
)

//go:embed wwwroot/css/*
//go:embed wwwroot/js/*
var staticFS embed.FS

func static() fs.FS {
	res, _ := fs.Sub(staticFS, "wwwroot")
	return res
}

func main() {
	port := ":7070"
	logger := slog.Default()
	logger.Info("Starting server", slog.String("port", port))

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(views.Home))

	fileHandler := http.StripPrefix("/assets/", http.FileServer(http.FS(static())))
	mux.Handle("/assets/", fileHandler)

	mux.Handle("/api/v1/user/", auth.Authenticate(users.UserMux()))

	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Server stopped")
}
