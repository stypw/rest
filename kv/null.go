package kv

type jsonNull struct {
	*jsonUndefined
}

var Null = (*jsonNull)(nil)

func NewNull() Element {
	return Null
}
func (elt *jsonNull) GetType() int {
	return NullType
}
func (elt *jsonNull) GetNull() interface{} {
	return nil
}

func (elt *jsonNull) TryGetNull() (interface{}, bool) {
	return nil, true
}

func (elt *jsonNull) GetValue() interface{} {
	return nil
}
func (elt *jsonNull) IsValue() bool {
	return true
}
func (elt *jsonNull) ToString() string {
	return "null"
}

func (elt *jsonNull) IsValid() bool {
	return true
}
