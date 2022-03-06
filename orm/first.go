package orm

import (
	"errors"
	"fmt"
	"rest/kv"
)

func readItem(fields []*field) (kv.Element, error) {

	var item kv.Element = kv.NewObject()
	for _, f := range fields {
		switch p := f.fieldValue.(type) {
		case *int64:
			item.Set(f.fieldName, kv.NewNumber(float64(*p)))
		case *string:
			item.Set(f.fieldName, kv.NewString(*p))
		case *float64:
			item.Set(f.fieldName, kv.NewNumber(*p))
		case *bool:
			item.Set(f.fieldName, kv.NewBoolean(*p))
		}
	}
	return item, nil
}

func (o *orm) First(where, order kv.Element) (kv.Element, error) {
	w, vs, err := parseAnd(where)
	if err != nil {
		return kv.Null, err
	}
	if w == "" {
		return kv.Null, errors.New("where can not empty")
	}
	od, err := parseOrder(order)
	if err != nil {
		return kv.Null, err
	}
	orderString := ""
	if od != "" {
		orderString = " order by " + od
	}

	sqlText := fmt.Sprintf("select * from %s where %s %s limit 0,1;", o.tableName, w, orderString)
	if rows, err := o.db.Query(sqlText, vs...); err == nil {
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
