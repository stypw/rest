package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	JSON "rest/json"
	"strings"

	orm "rest/orm"
)

type Router struct {
	Pattern string
	Db      *sql.DB
	Table   string
}

type sqlField struct {
	FieldType  string
	FieldName  string
	FieldValue interface{}
}

func makeFields(rows *sql.Rows) ([]*sqlField, []interface{}, error) {
	if cts, err := rows.ColumnTypes(); err == nil {
		var fields []*sqlField
		var pointers []interface{}
		for _, ct := range cts {
			field := &sqlField{FieldType: ct.DatabaseTypeName(), FieldName: ct.Name()}
			//include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL", "INT", and "BIGINT".
			switch ct.DatabaseTypeName() {
			case "INT", "BIGINT":
				{
					var v int64
					field.FieldValue = &v
					pointers = append(pointers, &v)
				}
			case "VARCHAR", "TEXT", "NVARCHAR":
				{
					var v string
					field.FieldValue = &v
					pointers = append(pointers, &v)
				}
			case "DECIMAL":
				{
					var v float64
					field.FieldValue = &v
					pointers = append(pointers, &v)
				}
			case "BOOL":
				{
					var v bool
					field.FieldValue = &v
					pointers = append(pointers, &v)
				}
			}
			fields = append(fields, field)
		}
		return fields, pointers, nil

	}
	return nil, nil, fmt.Errorf("error")
}

func boolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func writeItem(rows *sql.Rows, w http.ResponseWriter, fields []*sqlField, pointers []interface{}) {
	if err := rows.Scan(pointers...); err == nil {
		w.Write([]byte("{"))
		w.Write([]byte(`"code":0,"data": {`))
		first := true
		for _, f := range fields {
			if !first {
				w.Write([]byte(","))
			}
			first = false
			w.Write([]byte(fmt.Sprintf(`"%s":`, f.FieldName)))
			switch p := f.FieldValue.(type) {
			case *int64:
				w.Write([]byte(fmt.Sprintf("%d", *p)))
			case *string:
				w.Write([]byte(fmt.Sprintf(`"%s"`, *p)))
			case *float64:
				w.Write([]byte(fmt.Sprintf("%f", *p)))
			case *bool:
				w.Write([]byte(boolToString(*p)))
			}

		}
		w.Write([]byte("}"))
		w.Write([]byte("}"))
	} else {
		fmt.Println(err)
	}
}

func (r *Router) one(w http.ResponseWriter, req *http.Request, id string) {
	sqlText := fmt.Sprintf("select * from %s where id = ?;", r.Table)
	var iid int
	fmt.Fscanf(strings.NewReader(id), "%d", &iid)
	if rows, err := r.Db.Query(sqlText, iid); err == nil {
		defer rows.Close()
		if fields, pointers, err := makeFields(rows); err == nil {
			rows.Next()
			writeItem(rows, w, fields, pointers)
		}
		return
	}

	w.Write([]byte(fmt.Sprintf("Your Request{URL:%q,Method:GET,Id:%s}", req.URL, id)))
}

func (r *Router) list(w http.ResponseWriter, req *http.Request) {
	if body, err := JSON.FromStream(req.Body); err == nil {
		if where, y := body["where"]; y {
			if whereObj, y := where.(JSON.Object); y {
				if where3, ps, err := orm.Parse(whereObj); err == nil {
					whereString := " where " + where3
					sqlText := fmt.Sprintf("select * from %s %s;", r.Table, whereString)
					if rows, err := r.Db.Query(sqlText, ps...); err == nil {
						defer rows.Close()
						if fields, pointers, err := makeFields(rows); err == nil {
							for rows.Next() {
								writeItem(rows, w, fields, pointers)
							}
							return
						}
					}
				}
			}
		}
	}

	w.Write([]byte(`{"code":1,"error":"参数错误,body必须为正确的JSON且带有where字段"}`))

}

func (r *Router) search(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte(fmt.Sprintf("Your Request{URL:%q,Method:GET}", req.URL)))

}

var oneMatch = regexp.MustCompile(`/one/(\S+)$`)
var listMatch = regexp.MustCompile(`/list/?$`)
var searchMatch = regexp.MustCompile(`/search/?$`)

func (r *Router) get(w http.ResponseWriter, req *http.Request) {

	if result := oneMatch.FindStringSubmatch(req.URL.Path); len(result) > 1 {
		r.one(w, req, result[1])
		return
	}
	if listMatch.MatchString(req.URL.Path) {
		r.list(w, req)
		return
	}
	if searchMatch.MatchString(req.URL.Path) {
		r.search(w, req)
		return
	}
	w.Write([]byte(fmt.Sprintf("Your Request{URL:%q,Method:GET}", req.URL)))

}
func (r *Router) post(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Your Request{URL:%q,Method:POST}", req.URL)))
}
func (r *Router) put(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Your Request{URL:%q,Method:PUT}", req.URL)))
}
func (r *Router) delete(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Your Request{URL:%q,Method:DELETE}", req.URL)))
}

func (r *Router) all(w http.ResponseWriter, req *http.Request) {
	switch strings.ToLower(req.Method) {
	case "get":
		r.get(w, req)
	case "post":
		r.post(w, req)
	case "put":
		r.put(w, req)
	case "delete":
		r.delete(w, req)
	}
}

func (r *Router) Start() error {
	if r.Pattern == "" {
		return fmt.Errorf("PatternCannotEmpty")
	}
	http.HandleFunc(r.Pattern, r.all)
	if r.Pattern[len(r.Pattern)-1:] == "/" {
		http.HandleFunc(r.Pattern[:len(r.Pattern)-1], r.all)
	} else {
		http.HandleFunc(r.Pattern+"/", r.all)
	}
	return nil
}
