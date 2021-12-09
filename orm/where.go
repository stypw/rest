package orm

import (
	"fmt"
	JSON "rest/json"
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

type jsonBoolean bool

func (b jsonBoolean) toByte() byte {
	if b {
		return 1
	} else {
		return 0
	}
}

func GetJsonValue(v JSON.Value) interface{} {
	switch vv := v.(type) {
	case JSON.Null:
		return nil
	case JSON.Number:
		return vv
	case JSON.String:
		return vv
	case JSON.Boolean:
		return vv
	default:
		return nil
	}
}

func parseNotin(v JSON.Value) (string, []interface{}, error) {
	if array, y := v.(JSON.Array); y {
		if len(array) < 1 {
			return "empty-notin", nil, nil
		}
		var strs []string = make([]string, 0)
		var ponters []interface{} = make([]interface{}, 0)
		for _, item := range array {
			strs = append(strs, "?")
			ponters = append(ponters, GetJsonValue(item))
		}

		return fmt.Sprintf(" not in (%s)", strings.Join(strs, ",")), ponters, nil
	}
	return "valueinvalid-notin", nil, nil
}

func parseIn(v JSON.Value) (string, []interface{}, error) {
	if array, y := v.(JSON.Array); y {
		if len(array) < 1 {
			return "empty-in", nil, nil
		}
		var strs []string = make([]string, 0)
		var ponters []interface{} = make([]interface{}, 0)
		for _, item := range array {
			strs = append(strs, "?")
			ponters = append(ponters, GetJsonValue(item))
		}

		return fmt.Sprintf(" in (%s)", strings.Join(strs, ",")), ponters, nil
	}
	return "valueinvalid-in", nil, nil
}

func parseNot(v JSON.Value) (string, []interface{}, error) {
	return " <> ?", []interface{}{GetJsonValue(v)}, nil
}

func parseIs(v JSON.Value) (string, []interface{}, error) {
	return " = ?", []interface{}{GetJsonValue(v)}, nil
}

func parseLte(v JSON.Value) (string, []interface{}, error) {
	return " <= ?", []interface{}{GetJsonValue(v)}, nil
}

func parseLt(v JSON.Value) (string, []interface{}, error) {
	return " < ?", []interface{}{GetJsonValue(v)}, nil
}

func parseGte(v JSON.Value) (string, []interface{}, error) {
	return " >= ?", []interface{}{GetJsonValue(v)}, nil
}

func parseGt(v JSON.Value) (string, []interface{}, error) {
	return " > ?", []interface{}{GetJsonValue(v)}, nil
}

func parseLike(v JSON.Value) (string, []interface{}, error) {
	return " like ?", []interface{}{GetJsonValue(v)}, nil
}

func parseOr(v JSON.Value) (string, []interface{}, error) {
	var orList []string
	valid := false
	var params []interface{}
	if array, ok := v.(JSON.Array); ok {
		for _, item := range array {
			if obj, y := item.(JSON.Object); y {
				if strOr, ps, err := parseAnd(obj); err == nil {
					orList = append(orList, strOr)
					if len(ps) > 0 {
						params = append(params, ps...)
					}
					valid = true
				}
			}
		}
	}
	if valid {
		return fmt.Sprintf("(%s)", strings.Join(orList, " or ")), params, nil
	}
	return "", nil, fmt.Errorf("invalidor")
}

func parseDollar(k string, v JSON.Value) (string, []interface{}, error) {
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

func parseObject(obj JSON.Object) (string, []interface{}, error) {
	for k, v := range obj {
		if express, ps, err := parseDollar(k, v); err == nil {
			return express, ps, nil
		}
	}
	return "", nil, fmt.Errorf("invalidvalue")
}

func parseAnd(where JSON.Object) (string, []interface{}, error) {
	var ands []string
	var params []interface{}
	for k, v := range where {
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

		switch val := v.(type) {
		case JSON.Number:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, val)
		case JSON.String:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, val)
		case JSON.Boolean:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, jsonBoolean(val).toByte())
		case JSON.Object:
			if express, ps, err := parseObject(val); err == nil {
				ands = append(ands, fmt.Sprintf("%s %s", k, express))
				if len(ps) > 0 {
					params = append(params, ps...)
				}
			}
		}

	}
	return strings.Join(ands, " and "), params, nil
}
