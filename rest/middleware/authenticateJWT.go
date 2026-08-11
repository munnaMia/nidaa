package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"strings"

	"github.com/munnaMia/nidaa/util/conv"
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
		if len(parts) != 2 && parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		accessToken := parts[1]
		accTokenArr := strings.Split(accessToken, ".")
		if len(accTokenArr) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := accTokenArr[0]
		jwtBody := accTokenArr[1]
		jwtSignature := accTokenArr[2]

		massage := jwtHeader + "." + jwtBody
		scrBytes := []byte(mdlw.config.Service.SecretKey)
		msgBytes := []byte(massage)

		h := hmac.New(sha256.New, scrBytes)
		h.Write(msgBytes)
		newHash := h.Sum(nil)
		newSignature := conv.B64UrlEncoding(newHash)

		if jwtSignature != newSignature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
