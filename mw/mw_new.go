package mw

import (
	"net/http"
)

type MiddlewareFunc func(r *http.Request, msg *Message, extra Extra, next Next)

func (mw MiddlewareFunc) Call(r *http.Request, msg *Message, extra Extra, next Next) {
	mw(r, msg, extra, next)
}

func NewMiddleware(f MiddlewareFunc) Middleware {
	return MiddlewareFunc(f)
}
