package gn

func NewObject() Element {
	return &element{
		elementType: ObjectType,
		children:    make(map[string]Element),
	}
}

func (elt *element) GetProperty(k string) Element {
	if elt == nil {
		return null
	}
	if elt.elementType != ObjectType {
		return null
	}
	if elt.children == nil {
		return null
	}
	if e, y := elt.children[k]; y {
		return e
	}
	return null
}

func (elt *element) TryGetProperty(k string) (Element, bool) {
	if elt == nil {
		return null, false
	}
	if elt.elementType != ObjectType {
		return null, false
	}
	if elt.children == nil {
		return null, false
	}
	if e, y := elt.children[k]; y {
		return e, true
	}
	return null, false
}

func (elt *element) ObjectEnumerator() map[string]Element {
	if elt == nil {
		return nil
	}
	if elt.elementType != ObjectType {
		return nil
	}
	return elt.children
}

func (elt *element) Set(k string, v Element) Element {
	if elt == nil {
		return null
	}
	if elt.elementType != ObjectType {
		return null
	}
	if elt.children == nil {
		return null
	}
	elt.children[k] = v
	return elt
}
