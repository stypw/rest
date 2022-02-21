package gn

func NewString(v string) Element {
	return &element{
		elementType: StringType,
		stringValue: v,
	}
}
func (elt *element) GetString() string {
	if elt == nil {
		return ""
	}
	return elt.stringValue
}
func (elt *element) TryGetString() (string, bool) {
	if elt == nil {
		return "", false
	}
	if elt.elementType != StringType {
		return "", false
	}
	return elt.stringValue, true
}
