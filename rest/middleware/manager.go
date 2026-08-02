package middleware

import "net/http"

type mdlw func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []mdlw
}

// create a new middleware manager
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]mdlw, 0),
	}
}

// It register all the global middleware in an array.
func (mngr *Manager) GlobalMiddleware(middlewares ...mdlw) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

// It will wrap all the given local middlewares into a http handler,
// and it will follow the FIFO principle to wraps the middlewares
func (mngr *Manager) With(h http.Handler, localMiddlewares ...mdlw) http.Handler {
	handler := h

	for idx := len(localMiddlewares) - 1; idx >= 0; idx-- {
		handler = localMiddlewares[idx](handler)
	}

	return handler
}

// It will wrap all the given global middlewares into a http handler,
// and it will follow the FIFO principle to wraps the middlewares
func (mngr *Manager) Wrap(h http.Handler) http.Handler {
	handler := h

	for idx := len(mngr.globalMiddlewares) - 1; idx >= 0; idx-- {
		handler = mngr.globalMiddlewares[idx](handler)
	}

	return handler
}
