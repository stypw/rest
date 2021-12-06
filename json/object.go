package JSON

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Object map[string]Value

func (v Object) ToString() string {
	var vals []string
	for k, val := range v {
		vals = append(vals, fmt.Sprintf(`"%s":%s`, k, val.ToString()))
	}
	return fmt.Sprintf(`{%s}`, strings.Join(vals, ","))
}

func (v Object) parse(dec *json.Decoder) error {
	var currKey string
	var err error
	var interator = tokenIterator{
		dec: dec,
		booleanHandle: func(b bool) bool {
			if currKey == "" {
				err = errors.New("json fmt error")
				return false
			} else {
				v[currKey] = Boolean(b)
				currKey = ""
				return true
			}
		},
		stringHandle: func(s string) bool {
			if currKey == "" {
				currKey = s
				return true
			} else {
				v[currKey] = String(s)
				currKey = ""
				return true
			}
		},
		numberHandle: func(f float64) bool {
			if currKey == "" {
				err = errors.New("json fmt error")
				return false
			} else {
				v[currKey] = Number(f)
				currKey = ""
				return true
			}
		},
		nullHandle: func() bool {
			if currKey == "" {
				err = errors.New("json fmt error")
				return false
			} else {
				v[currKey] = Null()
				currKey = ""
				return true
			}
		},
		objectStartHandle: func() bool {
			if currKey == "" {
				err = errors.New("json fmt error")
				return false
			} else {
				var obj Object = make(Object)
				v[currKey] = obj
				obj.parse(dec)
				currKey = ""
				return true
			}
		},
		objectEndHandle: func() bool {
			return false
		},
		arrayStartHandle: func() bool {
			if currKey == "" {
				err = errors.New("json fmt error")
				return false
			} else {
				var arr Array = make(Array, 0)
				arr.parse(dec)
				v[currKey] = arr
				currKey = ""
				return true
			}
		},
		arrayEndHandle: func() bool {
			err = errors.New("json fmt error")
			return false
		},
	}

	interator.run()
	return err
}
