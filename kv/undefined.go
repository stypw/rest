package kv

type jsonUndefined struct{}

var undefined *jsonUndefined = (*jsonUndefined)(nil)

var Undefined = undefined

func (elt *jsonUndefined) GetType() int {
	return UndefinedType
}
func (elt *jsonUndefined) GetNull() interface{} {
	return nil
}
func (elt *jsonUndefined) TryGetNull() (interface{}, bool) {
	return nil, false
}
func (elt *jsonUndefined) GetBoolean() bool {
	return false
}
func (elt *jsonUndefined) TryGetBoolean() (bool, bool) {
	return false, false
}
func (elt *jsonUndefined) GetNumber() float64 {
	return 0
}
func (elt *jsonUndefined) TryGetNumber() (float64, bool) {
	return 0, false
}
func (elt *jsonUndefined) GetString() string {
	return ""
}
func (elt *jsonUndefined) TryGetString() (string, bool) {
	return "", false
}
func (elt *jsonUndefined) GetProperty(k string) Element {
	return undefined
}
func (elt *jsonUndefined) TryGetProperty(k string) (Element, bool) {
	return undefined, false
}
func (elt *jsonUndefined) GetElement(idx int) Element {
	return undefined
}
func (elt *jsonUndefined) TryGetElement(idx int) (Element, bool) {
	return undefined, false
}
func (elt *jsonUndefined) Set(k string, v Element) Element {
	return undefined
}
func (elt *jsonUndefined) Push(v Element) Element {
	return undefined
}
func (elt *jsonUndefined) Select(path string) Element {
	return undefined
}
func (elt *jsonUndefined) TrySelect(path string) (Element, bool) {
	return undefined, false
}
func (elt *jsonUndefined) ArrayEnumerator() []Element {
	return nil
}
func (elt *jsonUndefined) ObjectEnumerator() map[string]Element {
	return nil
}
func (elt *jsonUndefined) ToString() string {
	return "undefined"
}
func (elt *jsonUndefined) GetValue() interface{} {
	return nil
}
func (elt *jsonUndefined) IsValue() bool {
	return false
}

func (elt *jsonUndefined) IsValid() bool {
	return false
}
