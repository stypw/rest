package gn

func NewNumber(v float64) Element {
	return &element{
		elementType: NumberType,
		numberValue: v,
	}
}

func (elt *element) GetNumber() float64 {
	if elt == nil {
		return 0
	}
	return elt.numberValue
}

func (elt *element) TryGetNumber() (float64, bool) {
	if elt == nil {
		return 0, false
	}
	if elt.elementType != NumberType {
		return 0, false
	}
	return elt.numberValue, true
}
