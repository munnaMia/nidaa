package middleware

import "github.com/munnaMia/nidaa/internal/config"

type Middleware struct {
	config         *config.Configuration
	MaxBytesReader int64
}

// create a new middleware
func NewMiddleware(cnf *config.Configuration) *Middleware {
	return &Middleware{
		config: cnf,
	}
}
