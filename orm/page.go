package orm

import (
	"errors"
	"fmt"

	"github.com/stypw/rest/kv"
)

func (o *orm) Page(where, order kv.Element, page, size int) (kv.Element, error) {
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

	if page < 0 {
		page = 0
	}
	if size < 1 {
		size = 1
	}
	start := page * size
	end := (page + 1) * size

	sqlText := fmt.Sprintf("select * from %s where %s %s limit %d,%d;", o.tableName, w, orderString, start, end)
	if rows, err := o.db.Query(sqlText, vs...); err == nil {
		defer rows.Close()
		if fields, pointers, err := makeFields(rows); err == nil {
			array := kv.NewArray()
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
