package orm

import (
	"errors"
	"fmt"
	JSON "rest/json"
	"strings"
)

func parseCreate(item JSON.Object) ([]string, []string, []interface{}, error) {

	var fields []string = make([]string, 0)
	var marks []string = make([]string, 0)
	var values []interface{} = make([]interface{}, 0)

	for key, child := range item {
		switch ch := child.(type) {
		case *JSON.Null_T:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, nil)
			}
		case *JSON.Number_T:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, ch.Value)
			}
		case *JSON.String_T:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, ch.Value)
			}
		case *JSON.Boolean_T:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, ch.Value)
			}
		default:
			return nil, nil, nil, errors.New("unknowed data type")
		}
	}
	return fields, marks, values, nil
}

func (orm *Orm) Create(item JSON.Object) (*JSON.Number_T, error) {
	if item == nil {
		return JSON.Number(0), errors.New("item data can not empty")
	}
	fs, ms, vs, err := parseCreate(item)
	if err != nil {
		return JSON.Number(0), err
	}
	sqlText := fmt.Sprintf("insert into %s(%s) values(%s);", orm.TableName, strings.Join(fs, ","), strings.Join(ms, ","))
	ret, err := orm.Db.Exec(sqlText, vs...)
	if err != nil {
		return JSON.Number(0), err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return JSON.Number(0), err
	}
	return JSON.Number(float64(id)), nil
}
