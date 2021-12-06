package orm

import (
	"errors"
	"fmt"
	JSON "rest/json"
	"strings"
)

func parseSet(item JSON.Object) ([]string, []interface{}, error) {

	var sets []string = make([]string, 0)
	var values []interface{} = make([]interface{}, 0)

	for key, child := range item {
		switch ch := child.(type) {
		case *JSON.Null_T:
			{
				sets = append(sets, fmt.Sprintf("%s = ?"), key)
				values = append(values, nil)
			}
		case *JSON.Number_T:
			{
				sets = append(sets, fmt.Sprintf("%s = ?"), key)
				values = append(values, ch.Value)
			}
		case *JSON.String_T:
			{
				sets = append(sets, fmt.Sprintf("%s = ?"), key)
				values = append(values, ch.Value)
			}
		case *JSON.Boolean_T:
			{
				sets = append(sets, fmt.Sprintf("%s = ?"), key)
				values = append(values, ch.Value)
			}
		default:
			return nil, nil, errors.New("unknowed data type")
		}
	}
	return sets, values, nil
}

func (orm *Orm) Update(where JSON.Object, data JSON.Object) (*JSON.Number_T, error) {

	w, vs, err := parseAnd(where)
	if err != nil {
		return JSON.Number(0), err
	}
	if w == "" {
		return JSON.Number(0), errors.New("where can not empty")
	}
	sets, vals, err := parseSet(data)
	if err != nil {
		return JSON.Number(0), err
	}

	sqlText := fmt.Sprintf("update %s set %s where %s;", orm.TableName, strings.Join(sets, ","), w)
	values := append(vals, vs...)
	ret, err := orm.Db.Exec(sqlText, values...)
	if err != nil {
		return JSON.Number(0), err
	}
	id, err := ret.RowsAffected()
	if err != nil {
		return JSON.Number(0), nil
	}
	return JSON.Number(float64(id)), nil
}
