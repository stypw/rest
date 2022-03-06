package kv

import (
	"fmt"
	"strings"
)

type jsonObject struct {
	*jsonUndefined
	value map[string]Element
}

func NewObject() Element {
	return &jsonObject{
		value: make(map[string]Element),
	}
}

func (elt *jsonObject) GetType() int {
	return ObjectType
}

func (elt *jsonObject) GetProperty(k string) Element {
	if e, y := elt.value[k]; y {
		return e
	}
	return undefined
}

func (elt *jsonObject) TryGetProperty(k string) (Element, bool) {
	if e, y := elt.value[k]; y {
		return e, true
	}
	return undefined, false
}

func (elt *jsonObject) ObjectEnumerator() map[string]Element {
	return elt.value
}

func (elt *jsonObject) Set(k string, v Element) Element {
	elt.value[k] = v
	return elt
}
func (elt *jsonObject) ToString() string {
	childStrs := make([]string, 0)
	for k, c := range elt.ObjectEnumerator() {
		childStrs = append(childStrs, fmt.Sprintf("%q:%s", k, c.ToString()))
	}
	return "{" + strings.Join(childStrs, ",") + "}"
}

func (elt *jsonObject) Select(path string) Element {
	return Select(elt, path)
}
func (elt *jsonObject) TrySelect(path string) (Element, bool) {
	return TrySelect(elt, path)
}

func (elt *jsonObject) IsValid() bool {
	return true
}
