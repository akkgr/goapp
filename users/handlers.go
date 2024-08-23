package users

import "net/http"

func UserMux() http.Handler {
	userMux := http.NewServeMux()
	userMux.Handle("/profile", http.HandlerFunc(GetProfile))

	return http.StripPrefix("/api/v1/user", userMux)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(r.Context().Value("claims").([]byte))
}
