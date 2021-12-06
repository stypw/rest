package orm

import (
	"errors"
	"fmt"
	JSON "rest/json"
)

func readItem(fields []*field) (JSON.Object, error) {
	var item JSON.Object = make(JSON.Object)
	for _, f := range fields {
		switch p := f.FieldValue.(type) {
		case *int64:
			item[f.FieldName] = JSON.Number(float64(*p))
		case *string:
			item[f.FieldName] = JSON.String(*p)
		case *float64:
			item[f.FieldName] = JSON.Number(float64(*p))
		case *bool:
			item[f.FieldName] = JSON.Boolean(*p)
		}
	}
	return item, nil
}

func (orm *Orm) First(where, order JSON.Object) (JSON.Object, error) {
	w, vs, err := parseAnd(where)
	if err != nil {
		return nil, err
	}
	if w == "" {
		return nil, errors.New("where can not empty")
	}
	o, err := parseOrder(order)
	if err != nil {
		return nil, err
	}
	orderString := ""
	if o != "" {
		orderString = " order by " + o
	}

	sqlText := fmt.Sprintf("select * from %s where %s %s limit 0,1;", orm.TableName, w, orderString)
	if rows, err := orm.Db.Query(sqlText, vs...); err == nil {
		defer rows.Close()
		if fields, pointers, err := makeFields(rows); err == nil {
			if rows.Next() {
				if err := rows.Scan(pointers...); err == nil {
					return readItem(fields)
				} else {
					return nil, err
				}
			} else {
				return nil, nil
			}

		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
