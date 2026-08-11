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
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArray := strings.Split(header, " ") // extract the bearer and token
		if len(headerArray) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArray[1]
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
