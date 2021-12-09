package orm

import (
	"errors"
	"fmt"
	JSON "rest/json"
)

func (orm *Orm) Remove(where JSON.Object) (JSON.Number, error) {
	w, vs, err := parseAnd(where)
	if err != nil {
		return JSON.Number(0), err
	}
	if w == "" {
		return JSON.Number(0), errors.New("where can not empty")
	}

	sqlText := fmt.Sprintf("delete from %s where %s;", orm.TableName, w)
	ret, err := orm.Db.Exec(sqlText, vs...)
	if err != nil {
		return JSON.Number(0), err
	}
	id, err := ret.RowsAffected()
	if err != nil {
		return JSON.Number(0), nil
	}
	return JSON.Number(float64(id)), nil
}
