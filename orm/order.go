package orm

import (
	"errors"
	"fmt"
	"rest/gn"
	"strings"
)

func parseOrder(order gn.Element) (string, error) {
	if nil == order {
		return "", nil
	}
	orders := make([]string, 0)
	switch order.GetType() {
	case gn.ObjectType:
		{
			for key, child := range order.ObjectEnumerator() {
				switch child.GetType() {
				case gn.StringType:
					{
						ch := child.GetString()
						var od = strings.ToLower(strings.Trim(ch, ""))
						if od == "asc" || od == "desc" {
							orders = append(orders, fmt.Sprintf("%s %s", key, od))
							orders = append(orders, fmt.Sprintf("%s %s", key, od))

						} else {
							return "", errors.New("order fmt error")
						}
					}
				case gn.NumberType:
					{
						ch := child.GetNumber()
						if ch > 0 {
							orders = append(orders, key+" asc")
						} else {
							orders = append(orders, key+" desc")
						}

					}
				case gn.BooleanType:
					{
						ch := child.GetBoolean()
						if ch {
							orders = append(orders, key+" asc")
						} else {
							orders = append(orders, key+" desc")
						}
					}
				}
			}
		}
	case gn.ArrayType:
		{
			for _, item := range order.ArrayEnumerator() {
				if item.GetType() != gn.StringType {
					return "", errors.New("order fmt error")
				}
				orders = append(orders, item.GetString())
			}
		}
	case gn.NullType:
		{
			return "", nil
		}
	default:
		return "", errors.New("order fmt error")
	}

	if len(orders) > 0 {
		return strings.Join(orders, ","), nil
	}
	return "", nil
}
