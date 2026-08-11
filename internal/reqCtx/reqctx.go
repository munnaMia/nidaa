package reqctx

import "context"


type userCtxKey struct{}

// WithUser attaches the authenticated user to context
func WithUser(ctx context.Context)