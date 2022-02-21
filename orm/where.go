package orm

import (
	"fmt"
	"rest/gn"
	"strings"
)

/**********************************************************
$or,,$like,,$gt,,$gte,,$lt,,$lte,,$is,,$not,,$in,,$notin
{
	name:{$like:"l"},
	age:{$gt:20},
	nation:han,
	$or:[
		{skill:golang},
		{canwebiste:true}
	],
	class:{$in:[12,13,14]}
}
**********************************************************/

func chooseByte(b bool, t, f byte) byte {
	if b {
		return t
	}
	return f
}

func parseNotin(v gn.Element) (string, []interface{}, error) {
	if v.GetType() != gn.ArrayType {
		return "valueinvalid-notin", nil, nil
	}
	array := v.ArrayEnumerator()

	if len(array) < 1 {
		return "empty-notin", nil, nil
	}
	var strs []string = make([]string, 0)
	var ponters []interface{} = make([]interface{}, 0)
	for _, item := range array {
		if !item.IsValue() {
			continue
		}
		strs = append(strs, "?")
		ponters = append(ponters, item.RealValue())
	}

	return fmt.Sprintf(" not in (%s)", strings.Join(strs, ",")), ponters, nil
}

func parseIn(v gn.Element) (string, []interface{}, error) {

	if v.GetType() != gn.ArrayType {
		return "valueinvalid-in", nil, nil
	}
	array := v.ArrayEnumerator()

	if len(array) < 1 {
		return "empty-in", nil, nil
	}

	if len(array) < 1 {
		return "empty-in", nil, nil
	}
	var strs []string = make([]string, 0)
	var ponters []interface{} = make([]interface{}, 0)
	for _, item := range array {
		if !item.IsValue() {
			continue
		}
		strs = append(strs, "?")
		ponters = append(ponters, item.RealValue())
	}

	return fmt.Sprintf(" in (%s)", strings.Join(strs, ",")), ponters, nil

}

func parseNot(v gn.Element) (string, []interface{}, error) {
	return " <> ?", []interface{}{v.RealValue()}, nil
}

func parseIs(v gn.Element) (string, []interface{}, error) {
	return " = ?", []interface{}{v.RealValue()}, nil
}

func parseLte(v gn.Element) (string, []interface{}, error) {
	return " <= ?", []interface{}{v.RealValue()}, nil
}

func parseLt(v gn.Element) (string, []interface{}, error) {
	return " < ?", []interface{}{v.RealValue()}, nil
}

func parseGte(v gn.Element) (string, []interface{}, error) {
	return " >= ?", []interface{}{v.RealValue()}, nil
}

func parseGt(v gn.Element) (string, []interface{}, error) {
	return " > ?", []interface{}{v.RealValue()}, nil
}

func parseLike(v gn.Element) (string, []interface{}, error) {
	return " like ?", []interface{}{v.RealValue()}, nil
}

func parseOr(v gn.Element) (string, []interface{}, error) {
	var orList []string
	valid := false
	var params []interface{}
	array := v.ArrayEnumerator()
	for _, item := range array {
		if item.GetType() == gn.ObjectType {
			if strOr, ps, err := parseAnd(item); err == nil {
				orList = append(orList, strOr)
				if len(ps) > 0 {
					params = append(params, ps...)
				}
				valid = true
			}
		}
	}

	if valid {
		return fmt.Sprintf("(%s)", strings.Join(orList, " or ")), params, nil
	}
	return "", nil, fmt.Errorf("invalidor")
}

func parseDollar(k string, v gn.Element) (string, []interface{}, error) {
	switch k {
	case "$like":
		return parseLike(v)
	case "$gt":
		return parseGt(v)
	case "$gte":
		return parseGte(v)
	case "$lt":
		return parseLt(v)
	case "$lte":
		return parseLte(v)
	case "$is":
		return parseIs(v)
	case "$not":
		return parseNot(v)
	case "$in":
		return parseIn(v)
	case "$notin":
		return parseNotin(v)
	}
	return "", nil, fmt.Errorf("InvalidDollar")
}

func parseObject(obj gn.Element) (string, []interface{}, error) {
	for k, v := range obj.ObjectEnumerator() {
		if express, ps, err := parseDollar(k, v); err == nil {
			return express, ps, nil
		}
	}
	return "", nil, fmt.Errorf("invalidvalue")
}

func parseAnd(where gn.Element) (string, []interface{}, error) {
	var ands []string
	var params []interface{}
	for k, v := range where.ObjectEnumerator() {
		if k == "" {
			continue
		}

		if k[0] == '$' {
			if k == "$or" {
				if and, ps, err := parseOr(v); err == nil {
					ands = append(ands, and)
					if len(ps) > 0 {
						params = append(params, ps...)
					}
				}
			} else {
				continue
			}
		}

		switch v.GetType() {
		case gn.NumberType:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, v.RealValue())
		case gn.StringType:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, v.RealValue())
		case gn.BooleanType:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			val := v.GetBoolean()
			params = append(params, chooseByte(val, 1, 0))
		case gn.ObjectType:
			if express, ps, err := parseObject(v); err == nil {
				ands = append(ands, fmt.Sprintf("%s %s", k, express))
				if len(ps) > 0 {
					params = append(params, ps...)
				}
			}
		}
	}

	return strings.Join(ands, " and "), params, nil
}
