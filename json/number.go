package JSON

import "fmt"

type Number float64

func (v Number) ToString() string {
	return fmt.Sprintf("%g", v)
}
func GetNumber(v Value) float64 {
	if f, y := v.(Number); y {
		return float64(f)
	}
	return 0
}

func TryGetNumber(v Value) (float64, bool) {
	if f, y := v.(Number); y {
		return float64(f), true
	}
	return 0, false
}
