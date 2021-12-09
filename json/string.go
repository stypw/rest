package JSON

import "fmt"

type String string

func (v String) ToString() string {
	return fmt.Sprintf("\"%s\"", v)
}

func GetString(v Value) string {
	if s, y := v.(String); y {
		return string(s)
	}
	return ""
}

func TryGetString(v Value) (string, bool) {
	if s, y := v.(String); y {
		return string(s), true
	}
	return "", false
}
