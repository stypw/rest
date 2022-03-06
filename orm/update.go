package orm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/stypw/rest/kv"
)

func parseSet(item kv.Element) ([]string, []interface{}, error) {

	var sets []string = make([]string, 0)
	var values []interface{} = make([]interface{}, 0)
	if item.GetType() != kv.ObjectType {
		return nil, nil, errors.New("unknowed data type")
	}
	for key, child := range item.ObjectEnumerator() {
		switch child.GetType() {
		case kv.NullType:
			{
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, nil)
			}
		case kv.NumberType:
			{
				//ch, _ := kv.ToNumber(child)
				ch := child.GetNumber()
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, ch)
			}
		case kv.StringType:
			{
				// ch, _ := kv.ToString(child)
				ch := child.GetString()
				sets = append(sets, fmt.Sprintf("%s = ?", key))
				values = append(values, ch)
			}
		case kv.BooleanType:
			{
				// ch, _ := kv.ToBoolean(child)
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

func (o *orm) Update(where kv.Element, data kv.Element) (int64, error) {

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
