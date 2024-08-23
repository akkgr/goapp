package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		bearer := strings.Split(token, " ")
		if len(bearer) != 2 && bearer[0] != "Bearer" {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		value, err := jwt.ParseWithClaims(
			bearer[1],
			&UserClaims{},
			func(*jwt.Token) (interface{}, error) {
				return Salt, nil
			})
		if err != nil {
			http.Error(w, fmt.Sprintf("%v. %v", http.StatusText(http.StatusForbidden), err), http.StatusForbidden)
			return
		}

		claims, ok := value.Claims.(*UserClaims)
		if !(ok && value.Valid) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), claimsKey, claims))

		next.ServeHTTP(w, r)
	})
}
