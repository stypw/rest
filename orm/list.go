package orm

import (
	"errors"
	"fmt"
	"rest/gn"
)

func (o *orm) List(where, order gn.Element) (gn.Element, error) {
	w, vs, err := parseAnd(where)
	if err != nil {
		return nil, err
	}
	if w == "" {
		return nil, errors.New("where can not empty")
	}
	od, err := parseOrder(order)
	if err != nil {
		return nil, err
	}
	orderString := ""
	if od != "" {
		orderString = " order by " + od
	}

	sqlText := fmt.Sprintf("select * from %s where %s %s;", o.tableName, w, orderString)
	if rows, err := o.db.Query(sqlText, vs...); err == nil {
		defer rows.Close()
		if fields, pointers, err := makeFields(rows); err == nil {
			array := gn.NewArray()
			for rows.Next() {
				if err := rows.Scan(pointers...); err == nil {
					item, err := readItem(fields)
					if err != nil {
						return nil, err
					}
					array.Push(item)
				} else {
					return nil, err
				}
			}
			return array, nil

		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
