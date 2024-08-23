package main

import (
	"goapp/auth"
	"goapp/users"
	"log/slog"
	"net/http"
)

func main() {
	logger := slog.Default()
	logger.Info("Starting server")

	mux := http.NewServeMux()

	mux.Handle("/api/v1/user/", auth.Authenticate(users.UserMux()))

	http.ListenAndServe(":8080", mux)
	logger.Info("Server stopped")
}
