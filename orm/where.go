package orm

import (
	"fmt"
	"strings"

	"github.com/stypw/rest/kv"
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

func parseNotin(v kv.Element) (string, []interface{}, error) {
	if v.GetType() != kv.ArrayType {
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
		ponters = append(ponters, item.GetValue())
	}

	return fmt.Sprintf(" not in (%s)", strings.Join(strs, ",")), ponters, nil
}

func parseIn(v kv.Element) (string, []interface{}, error) {

	if v.GetType() != kv.ArrayType {
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
		ponters = append(ponters, item.GetValue())
	}

	return fmt.Sprintf(" in (%s)", strings.Join(strs, ",")), ponters, nil

}

func parseNot(v kv.Element) (string, []interface{}, error) {
	return " <> ?", []interface{}{v.GetValue()}, nil
}

func parseIs(v kv.Element) (string, []interface{}, error) {
	return " = ?", []interface{}{v.GetValue()}, nil
}

func parseLte(v kv.Element) (string, []interface{}, error) {
	return " <= ?", []interface{}{v.GetValue()}, nil
}

func parseLt(v kv.Element) (string, []interface{}, error) {
	return " < ?", []interface{}{v.GetValue()}, nil
}

func parseGte(v kv.Element) (string, []interface{}, error) {
	return " >= ?", []interface{}{v.GetValue()}, nil
}

func parseGt(v kv.Element) (string, []interface{}, error) {
	return " > ?", []interface{}{v.GetValue()}, nil
}

func parseLike(v kv.Element) (string, []interface{}, error) {
	return " like ?", []interface{}{v.GetValue()}, nil
}

func parseOr(v kv.Element) (string, []interface{}, error) {
	var orList []string
	valid := false
	var params []interface{}
	array := v.ArrayEnumerator()
	for _, item := range array {
		if item.GetType() == kv.ObjectType {
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

func parseDollar(k string, v kv.Element) (string, []interface{}, error) {
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

func parseObject(obj kv.Element) (string, []interface{}, error) {
	for k, v := range obj.ObjectEnumerator() {
		if express, ps, err := parseDollar(k, v); err == nil {
			return express, ps, nil
		}
	}
	return "", nil, fmt.Errorf("invalidvalue")
}

func parseAnd(where kv.Element) (string, []interface{}, error) {
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
		case kv.NumberType:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, v.GetValue())
		case kv.StringType:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			params = append(params, v.GetValue())
		case kv.BooleanType:
			ands = append(ands, fmt.Sprintf("%s = ?", k))
			val := v.GetBoolean()
			params = append(params, chooseByte(val, 1, 0))
		case kv.ObjectType:
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
