package middleware

type Middleware struct {
	MaxBytesReader int64
}

// create a new middleware
func NewMiddleware() *Middleware {
	return &Middleware{}
}
