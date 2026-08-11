package reqctx

import (
	"context"
	"strconv"

	"github.com/munnaMia/nidaa/internal/domain"
)

type jwtCtxKey struct{}

// WithUser attaches the authenticated user to context
func WithUser(ctx context.Context, payload *domain.Payload) context.Context {
	return context.WithValue(ctx, jwtCtxKey{}, payload)

}

// GetUserPayload extracts the UserPayload from context
func GetJWTPayload(ctx context.Context) (*domain.Payload, bool) {
	jwtP, ok := ctx.Value(jwtCtxKey{}).(*domain.Payload)
	return jwtP, ok
}

// GetUserID is a fast convenience helper to fetch just the User ID
func GetUserID(ctx context.Context) (int64, bool) {
	jwtP, ok := GetJWTPayload(ctx)
	if !ok || jwtP == nil {
		return 0, false
	}

	id, err := strconv.ParseInt(jwtP.Sub, 10, 64)
	if err != nil {
		return 0, false
	}
	
	return id, true
}
