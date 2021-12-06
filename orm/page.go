package orm

import (
	"errors"
	"fmt"
	JSON "rest/json"
)

func (orm *Orm) Page(where, order JSON.Object, page, size int) (JSON.Array, error) {
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

	if page < 0 {
		page = 0
	}
	if size < 1 {
		size = 1
	}
	start := page * size
	end := (page + 1) * size

	sqlText := fmt.Sprintf("select * from %s where %s %s limit %d,%d;", orm.TableName, w, orderString, start, end)
	if rows, err := orm.Db.Query(sqlText, vs...); err == nil {
		defer rows.Close()
		if fields, pointers, err := makeFields(rows); err == nil {
			array := make(JSON.Array, 0)
			for rows.Next() {
				if err := rows.Scan(pointers...); err == nil {
					item, err := readItem(fields)
					if err != nil {
						return nil, err
					}
					array = append(array, item)
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
