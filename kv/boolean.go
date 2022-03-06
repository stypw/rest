package kv

type jsonBoolean struct {
	*jsonUndefined
	value bool
}

func NewBoolean(v bool) Element {
	return &jsonBoolean{
		value: v,
	}
}
func (elt *jsonBoolean) GetType() int {
	return BooleanType
}

func (elt *jsonBoolean) GetBoolean() bool {
	return elt.value
}

func (elt *jsonBoolean) TryGetBoolean() (bool, bool) {
	return elt.value, true
}

func (elt *jsonBoolean) GetValue() interface{} {
	return elt.value
}
func (elt *jsonBoolean) IsValue() bool {
	return true
}

func (elt *jsonBoolean) ToString() string {
	if elt.value {
		return "true"
	}
	return "false"
}

func (elt *jsonBoolean) IsValid() bool {
	return true
}
