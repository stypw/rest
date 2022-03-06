package kv

import "fmt"

type jsonNumber struct {
	*jsonUndefined
	value float64
}

func NewNumber(v float64) Element {
	return &jsonNumber{
		value: v,
	}
}
func (elt *jsonNumber) GetType() int {
	return NumberType
}
func (elt *jsonNumber) GetNumber() float64 {
	return elt.value
}

func (elt *jsonNumber) TryGetNumber() (float64, bool) {
	return elt.value, true
}

func (elt *jsonNumber) GetValue() interface{} {
	return elt.value
}
func (elt *jsonNumber) IsValue() bool {
	return true
}
func (elt *jsonNumber) ToString() string {
	return fmt.Sprintf("%g", elt.value)
}

func (elt *jsonNumber) IsValid() bool {
	return true
}
