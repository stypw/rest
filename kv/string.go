package kv

import "fmt"

type jsonString struct {
	*jsonUndefined
	value string
}

func NewString(v string) Element {
	return &jsonString{
		value: v,
	}
}
func (elt *jsonString) GetType() int {
	return StringType
}

func (elt *jsonString) GetString() string {
	return elt.value
}
func (elt *jsonString) TryGetString() (string, bool) {
	return elt.value, true
}

func (elt *jsonString) GetValue() interface{} {
	return elt.value
}
func (elt *jsonString) IsValue() bool {
	return true
}

func (elt *jsonString) ToString() string {
	return fmt.Sprintf("%q", elt.value)
}

func (elt *jsonString) IsValid() bool {
	return true
}
