package middlewares

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/utils"
	"net/http"
	"time"
)

func (m middleware) WithAuth() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			authClaims, err := utils.JwtValidate(auth)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid token %v", err), http.StatusForbidden)
				return
			}

			if authClaims.StandardClaims.ExpiresAt < time.Now().Unix() {

			}

			ctx := context.WithValue(r.Context(), "auth", authClaims)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
