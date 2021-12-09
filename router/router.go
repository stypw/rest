package router

import (
	"database/sql"
	"fmt"
	"net/http"
	JSON "rest/json"
	"rest/orm"
)

type RestHandle func(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error)

type Router struct {
	Pattern string
	Table   string
	Db      *sql.DB
	Before  RestHandle
	orm     *orm.Orm
}

func (r *Router) onRequest(handle RestHandle) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var responseData JSON.Object = make(JSON.Object)
		before := r.Before
		if before != nil {
			code, _, err := before(w, req)
			if code != 0 {
				responseData["code"] = JSON.Number(float64(code))
				responseData["error"] = JSON.String(err.Error())
				w.Write([]byte(responseData.ToString()))
				return
			}
		}

		if handle != nil {
			code, data, err := handle(w, req)
			if code == 0 {
				responseData["code"] = JSON.Number(float64(code))
				responseData["data"] = data
			} else {
				responseData["code"] = JSON.Number(float64(code))
				responseData["error"] = JSON.String(err.Error())
			}
			w.Write([]byte(responseData.ToString()))
		}

	}
}

func (r *Router) Start(db *sql.DB) error {
	r.Db = db
	if r.Pattern == "" {
		return fmt.Errorf("PatternCannotEmpty")
	}

	pattern := r.Pattern
	if pattern[len(r.Pattern)-1:] != "/" {
		pattern = pattern + "/"
	}

	r.orm = &orm.Orm{Db: r.Db, TableName: r.Table}

	http.HandleFunc(pattern+"create", r.onRequest(r.create))
	http.HandleFunc(pattern+"create/", r.onRequest(r.create))

	http.HandleFunc(pattern+"remove", r.onRequest(r.remove))
	http.HandleFunc(pattern+"remove/", r.onRequest(r.remove))

	http.HandleFunc(pattern+"update", r.onRequest(r.update))
	http.HandleFunc(pattern+"update/", r.onRequest(r.update))

	http.HandleFunc(pattern+"first", r.onRequest(r.first))
	http.HandleFunc(pattern+"first/", r.onRequest(r.first))

	http.HandleFunc(pattern+"list", r.onRequest(r.list))
	http.HandleFunc(pattern+"list/", r.onRequest(r.list))

	http.HandleFunc(pattern+"page", r.onRequest(r.page))
	http.HandleFunc(pattern+"page/", r.onRequest(r.page))

	return nil
}
