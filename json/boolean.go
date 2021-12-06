package JSON

type Boolean_T struct{ Value bool }

func Boolean(v bool) *Boolean_T {
	return &Boolean_T{Value: v}
}

func (v *Boolean_T) toString() string {
	if v.Value {
		return "true"
	} else {
		return "false"
	}
}
