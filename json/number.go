package JSON

import "fmt"

type Number_T struct {
	Value float64
}

func Number(v float64) *Number_T {
	return &Number_T{Value: v}
}

func (v *Number_T) ToString() string {
	return fmt.Sprintf("%f", float64(v.Value))
}
