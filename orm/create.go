package orm

import (
	"errors"
	"fmt"
	"rest/gn"
	"strings"
)

func parseCreate(item gn.Element) ([]string, []string, []interface{}, error) {

	var fields []string = make([]string, 0)
	var marks []string = make([]string, 0)
	var values []interface{} = make([]interface{}, 0)
	if item.GetType() != gn.ObjectType {
		return nil, nil, nil, errors.New("unknowed data type")
	}
	for key, child := range item.ObjectEnumerator() {
		switch child.GetType() {
		case gn.NullType:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, nil)
			}
		case gn.NumberType:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, child.GetNumber())
			}
		case gn.StringType:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, child.GetString())
			}
		case gn.BooleanType:
			{
				fields = append(fields, key)
				marks = append(marks, "?")
				values = append(values, child.GetBoolean())
			}
		default:
			return nil, nil, nil, errors.New("unknowed data type")
		}
	}
	return fields, marks, values, nil
}

func (o *orm) Create(item gn.Element) (int64, error) {
	if item == nil {
		return 0, errors.New("item data can not empty")
	}
	fs, ms, vs, err := parseCreate(item)
	if err != nil {
		return 0, err
	}
	sqlText := fmt.Sprintf("insert into %s(%s) values(%s);", o.tableName, strings.Join(fs, ","), strings.Join(ms, ","))
	ret, err := o.db.Exec(sqlText, vs...)
	if err != nil {
		return 0, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
