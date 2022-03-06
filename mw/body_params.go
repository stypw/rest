package mw

import (
	"net/http"

	"github.com/stypw/rest/df"
	"github.com/stypw/rest/kv"
)

var JsonBodyMiddleware = NewMiddleware(func(r *http.Request, msg *Message, extra Extra, next Next) {
	body := kv.FromStream(r.Body)
	if body == nil {
		msg.Code = df.HTTP_STATUS_PARAM_ERROR
		msg.Msg = "body为空"
		return
	}
	extra["body"] = body
	next(r, msg, extra)
})
