package gn

const (
	UndefinedType = iota
	NullType
	BooleanType
	NumberType
	StringType
	ObjectType
	ArrayType
)

type Element interface {
	GetType() int

	GetNull() interface{}
	TryGetNull() (interface{}, bool)

	GetBoolean() bool
	TryGetBoolean() (bool, bool)

	GetNumber() float64
	TryGetNumber() (float64, bool)

	GetString() string
	TryGetString() (string, bool)

	GetProperty(k string) Element
	TryGetProperty(k string) (Element, bool)

	GetElement(idx int) Element
	TryGetElement(idx int) (Element, bool)

	Set(k string, v Element) Element
	Push(v Element) Element

	Select(path string) Element
	TrySelect(path string) (Element, bool)

	ArrayEnumerator() []Element
	ObjectEnumerator() map[string]Element

	ToString() string

	RealValue() interface{}

	IsValue() bool
}

type element struct {
	elementType int
	children    map[string]Element
	arrays      []Element
	nullValue   interface{}
	boolValue   bool
	numberValue float64
	stringValue string
}

var null *element = nil

var Null Element = null

func (elt *element) GetType() int {
	if elt == nil {
		return -1
	}
	return elt.elementType
}

func ToElement(s interface{}) Element {
	if e, y := s.(*element); y {
		return e
	}
	return null
}
