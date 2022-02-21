package gn

func NewArray() Element {
	return &element{
		elementType: ArrayType,
		arrays:      make([]Element, 0),
	}
}

func (elt *element) GetElement(idx int) Element {
	if elt == nil {
		return null
	}
	if elt.elementType != ArrayType {
		return null
	}
	if elt.arrays == nil {
		return null
	}
	l := len(elt.arrays)
	if idx >= 0 && idx < l {
		return elt.arrays[idx]
	}
	return null
}

func (elt *element) TryGetElement(idx int) (Element, bool) {
	if elt == nil {
		return null, false
	}
	if elt.elementType != ArrayType {
		return null, false
	}
	if elt.arrays == nil {
		return null, false
	}
	l := len(elt.arrays)
	if idx >= 0 && idx < l {
		return elt.arrays[idx], true
	}
	return null, false
}

func (elt *element) ArrayEnumerator() []Element {
	if elt == nil {
		return nil
	}
	if elt.elementType != ArrayType {
		return nil
	}
	return elt.arrays
}

func (elt *element) Push(item Element) Element {
	if elt == nil {
		return null
	}
	if elt.elementType != ArrayType {
		return null
	}
	elt.arrays = append(elt.arrays, item)
	return elt
}
