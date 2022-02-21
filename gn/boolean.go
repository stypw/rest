package gn

func NewBoolean(v bool) Element {
	return &element{
		elementType: BooleanType,
		boolValue:   v,
	}
}

func (elt *element) GetBoolean() bool {
	if elt == nil {
		return false
	}
	return elt.boolValue
}

func (elt *element) TryGetBoolean() (bool, bool) {
	if elt == nil {
		return false, false
	}
	if elt.elementType != BooleanType {
		return false, false
	}
	return elt.boolValue, true
}
