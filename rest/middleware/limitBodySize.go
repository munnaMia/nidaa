package middleware

import "net/http"

// limit req body size.
func (mdlw *Middleware) LimitBodySize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, mdlw.MaxBytesReader)
		next.ServeHTTP(w,r)
	})
}
