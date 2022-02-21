package rt

import "rest/mw"

func NewRouter() Router {
	return &router{middlewares: make([]mw.Middleware, 0), middleIndex: 0}
}
