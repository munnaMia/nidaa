package middleware

import (
	"github.com/munnaMia/nidaa/internal/config"
	"github.com/munnaMia/nidaa/internal/domain"
)

type Middleware struct {
	config         *config.Configuration
	tkSvr          domain.TokenService
	MaxBytesReader int64
}

// create a new middleware
func NewMiddleware(cnf *config.Configuration, tkSvr domain.TokenService) *Middleware {
	return &Middleware{
		config: cnf,
		tkSvr:  tkSvr,
	}
}
