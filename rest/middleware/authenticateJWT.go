package middleware

import (
	"net/http"
	"strings"

	reqctx "github.com/munnaMia/nidaa/internal/reqCtx"
)

// validate a jwt token for authorize users
func (mdlw *Middleware) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ") // extract the bearer and token
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		accessToken := parts[1]

		payload, err := mdlw.tkSvr.ValidateToken(accessToken)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(reqctx.WithUser(r.Context(), payload)))
	})
}
