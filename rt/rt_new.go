package rt

import "github.com/stypw/rest/mw"

func NewRouter() Router {
	return &router{middlewares: make([]mw.Middleware, 0), middleIndex: 0}
}
