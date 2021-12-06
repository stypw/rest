package router

import (
	"database/sql"
	"fmt"
	"net/http"
	JSON "rest/json"
	"rest/orm"
)

type Router struct {
	Pattern string
	Table   string
	Db      *sql.DB

	orm *orm.Orm
}

type restHandle func(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error)

func (r *Router) all(handle restHandle) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var responseData JSON.Object = make(JSON.Object)
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

func (r *Router) Start() error {
	if r.Pattern == "" {
		return fmt.Errorf("PatternCannotEmpty")
	}

	pattern := r.Pattern
	if pattern[len(r.Pattern)-1:] != "/" {
		pattern = pattern + "/"
	}

	r.orm = &orm.Orm{Db: r.Db, TableName: r.Table}

	http.HandleFunc(pattern+"create", r.all(r.create))
	http.HandleFunc(pattern+"create/", r.all(r.create))

	http.HandleFunc(pattern+"remove", r.all(r.remove))
	http.HandleFunc(pattern+"remove/", r.all(r.remove))

	http.HandleFunc(pattern+"update", r.all(r.update))
	http.HandleFunc(pattern+"update/", r.all(r.update))

	http.HandleFunc(pattern+"first", r.all(r.first))
	http.HandleFunc(pattern+"first/", r.all(r.first))

	http.HandleFunc(pattern+"list", r.all(r.list))
	http.HandleFunc(pattern+"list/", r.all(r.list))

	http.HandleFunc(pattern+"page", r.all(r.page))
	http.HandleFunc(pattern+"page/", r.all(r.page))

	return nil
}
