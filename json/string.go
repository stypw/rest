package JSON

import "fmt"

type String_T struct{ Value string }

func String(v string) *String_T {
	return &String_T{Value: v}
}
func (v *String_T) toString() string {
	return fmt.Sprintf("\"%s\"", v.Value)
}

func AsString(v Value) (string, bool) {
	if s, y := v.(*String_T); y {
		return s.Value, true
	}
	return "", false
}
