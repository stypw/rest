package rt

import (
	"fmt"
	"net/http"
	"rest/mw"
	"time"
)

type Router interface {
	UseMiddleware(mw mw.Middleware)
	onRequest(r *http.Request, msg *mw.Message, extra mw.Extra)
}

type routerRecord struct {
	pattern string
	factory RouterFactory
}

var routers []*routerRecord = make([]*routerRecord, 0)

func RegisterRouter(pattern string, factory RouterFactory) {
	routers = append(routers, &routerRecord{pattern: pattern, factory: factory})
}

func doRouter(record *routerRecord) {
	http.HandleFunc(record.pattern, func(rw http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		router := record.factory.Create()
		msg := &mw.Message{}
		extra := make(map[string]interface{})
		router.onRequest(r, msg, extra)
		rw.WriteHeader(msg.Code)
		rw.Write([]byte(msg.Msg))
		fmt.Printf("%d毫秒:%s\n", time.Since(timeStart)/time.Millisecond, record.pattern)
	})
}

func Start() {
	for _, r := range routers {
		doRouter(r)
		fmt.Println(r.pattern)
	}
}
