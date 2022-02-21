package gn

func NewNull() Element {
	return &element{
		elementType: NullType,
		nullValue:   nil,
	}
}

func (elt *element) GetNull() interface{} {
	if elt == nil {
		return nil
	}
	return elt.nullValue
}

func (elt *element) TryGetNull() (interface{}, bool) {
	if elt == nil {
		return nil, false
	}
	if elt.elementType != NullType {
		return nil, false
	}
	return elt.nullValue, true
}
