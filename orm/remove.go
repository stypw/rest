package orm

import (
	"errors"
	"fmt"
	"rest/kv"
)

func (o *orm) Remove(where kv.Element) (int64, error) {
	w, vs, err := parseAnd(where)
	if err != nil {
		return 0, err
	}
	if w == "" {
		return 0, errors.New("where can not empty")
	}

	sqlText := fmt.Sprintf("delete from %s where %s;", o.tableName, w)
	ret, err := o.db.Exec(sqlText, vs...)
	if err != nil {
		return 0, err
	}
	id, err := ret.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
