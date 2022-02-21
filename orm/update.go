package orm

import (
	"errors"
	"fmt"
	"rest/gn"
	"strings"
)

func parseSet(item gn.Element) ([]string, []interface{}, error) {

	var sets []string = make([]string, 0)
	var values []interface{} = make([]interface{}, 0)
	if item.GetType() != gn.ObjectType {
		return nil, nil, errors.New("unknowed data type")
	}
	for key, child := range item.ObjectEnumerator() {
		switch child.GetType() {
		case gn.NullType:
			{
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, nil)
			}
		case gn.NumberType:
			{
				//ch, _ := gn.ToNumber(child)
				ch := child.GetNumber()
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, ch)
			}
		case gn.StringType:
			{
				// ch, _ := gn.ToString(child)
				ch := child.GetString()
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, ch)
			}
		case gn.BooleanType:
			{
				// ch, _ := gn.ToBoolean(child)
				ch := child.GetBoolean()
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, ch)
			}
		default:
			return nil, nil, errors.New("unknowed data type")
		}
	}
	return sets, values, nil
}

func (o *orm) Update(where gn.Element, data gn.Element) (int64, error) {

	w, vs, err := parseAnd(where)
	if err != nil {
		return 0, err
	}
	if w == "" {
		return 0, errors.New("where can not empty")
	}
	sets, vals, err := parseSet(data)
	if err != nil {
		return 0, err
	}

	sqlText := fmt.Sprintf("update %s set %s where %s;", o.tableName, strings.Join(sets, ","), w)
	values := append(vals, vs...)
	ret, err := o.db.Exec(sqlText, values...)
	if err != nil {
		return 0, err
	}
	id, err := ret.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
