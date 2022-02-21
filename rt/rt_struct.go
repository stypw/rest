package rt

import (
	"net/http"
	"rest/mw"
)

type router struct {
	middlewares []mw.Middleware
	middleIndex int
}

func (rter *router) callMiddleware(r *http.Request, msg *mw.Message, extra mw.Extra) {
	l := len(rter.middlewares)
	if rter.middleIndex >= l {
		return
	}
	m := rter.middlewares[rter.middleIndex]
	rter.middleIndex++
	m.Call(r, msg, extra, rter.callMiddleware)
}
func (rter *router) onRequest(r *http.Request, msg *mw.Message, extra mw.Extra) {
	rter.middleIndex = 0
	rter.callMiddleware(r, msg, extra)
}

func (rter *router) UseMiddleware(middleware mw.Middleware) {
	rter.middlewares = append(rter.middlewares, middleware)
}
