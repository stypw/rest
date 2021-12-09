package JSON

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Array []Value

func (v Array) ToString() string {
	var vals []string
	for _, val := range v {
		vals = append(vals, val.ToString())
	}
	return fmt.Sprintf(`[%s]`, strings.Join(vals, ","))
}

func GetItem(v Value, i int) Value {
	if v == nil || i < 0 {
		return nil
	}
	if a, y := v.(Array); y {
		if i >= len(a) {
			return nil
		}
		return a[i]
	}
	return nil
}

func TryGetItem(v Value, i int) (Value, bool) {
	if v == nil || i < 0 {
		return nil, false
	}
	if a, y := v.(Array); y {
		if i >= len(a) {
			return nil, false
		}
		return a[i], true
	}
	return nil, false
}

func (v *Array) parse(dec *json.Decoder) error {
	var err error
	var interator = tokenIterator{
		dec: dec,
		booleanHandle: func(b bool) bool {
			*v = append(*v, Boolean(b))
			return true
		},
		stringHandle: func(s string) bool {
			*v = append(*v, String(s))
			return true
		},
		numberHandle: func(f float64) bool {
			*v = append(*v, Number(f))
			return true
		},
		nullHandle: func() bool {
			*v = append(*v, Null{})
			return true
		},
		objectStartHandle: func() bool {
			var obj = Object{}
			er := obj.parse(dec)
			if er != nil {
				err = er
				return false
			}
			*v = append(*v, obj)
			return true
		},
		objectEndHandle: func() bool {
			err = errors.New("json fmt error")
			return false
		},
		arrayStartHandle: func() bool {
			var arr Array = make(Array, 0)
			er := arr.parse(dec)
			if er != nil {
				err = er
				return false
			}
			*v = append(*v, arr)
			return true
		},
		arrayEndHandle: func() bool {
			return false
		},
	}

	interator.run()
	return err
}
