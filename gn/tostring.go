package gn

import (
	"fmt"
	"strings"
)

func (elt *element) ToString() string {
	if elt == nil {
		return ""
	}
	switch elt.elementType {
	case NullType:
		{
			return "null"
		}
	case BooleanType:
		{
			if elt.boolValue {
				return "true"
			}
			return "false"
		}
	case NumberType:
		{
			return fmt.Sprintf("%g", elt.numberValue)
		}
	case StringType:
		{
			return fmt.Sprintf("%q", elt.stringValue)
		}
	case ObjectType:
		{
			childStrs := make([]string, 0)
			for k, c := range elt.ObjectEnumerator() {
				childStrs = append(childStrs, fmt.Sprintf("%q:%s", k, c.ToString()))
			}
			return "{" + strings.Join(childStrs, ",") + "}"
		}
	case ArrayType:
		{
			childStrs := make([]string, 0)
			for _, i := range elt.ArrayEnumerator() {
				childStrs = append(childStrs, i.ToString())
			}
			return "[" + strings.Join(childStrs, ",") + "]"
		}
	}
	return ""
}
