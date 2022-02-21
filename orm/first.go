package orm

import (
	"errors"
	"fmt"
	"rest/gn"
)

func readItem(fields []*field) (gn.Element, error) {

	var item gn.Element = gn.NewObject()
	for _, f := range fields {
		switch p := f.fieldValue.(type) {
		case *int64:
			item.Set(f.fieldName, gn.NewNumber(float64(*p)))
		case *string:
			item.Set(f.fieldName, gn.NewString(*p))
		case *float64:
			item.Set(f.fieldName, gn.NewNumber(*p))
		case *bool:
			item.Set(f.fieldName, gn.NewBoolean(*p))
		}
	}
	return item, nil
}

func (o *orm) First(where, order gn.Element) (gn.Element, error) {
	w, vs, err := parseAnd(where)
	if err != nil {
		return gn.Null, err
	}
	if w == "" {
		return gn.Null, errors.New("where can not empty")
	}
	od, err := parseOrder(order)
	if err != nil {
		return gn.Null, err
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
