package orm

import (
	"errors"
	"fmt"
	JSON "rest/json"
	"strings"
)

func parseOrder(order JSON.Value) (string, error) {
	if nil == order {
		return "", nil
	}
	orders := make([]string, 0)
	switch vd := order.(type) {
	case JSON.Object:
		{
			for key, child := range vd {
				switch ch := child.(type) {
				case *JSON.String_T:
					{
						var od = strings.ToLower(strings.Trim(ch.Value, ""))
						if od == "asc" || od == "desc" {
							orders = append(orders, fmt.Sprintf("%s %s", key, od))
							orders = append(orders, fmt.Sprintf("%s %s", key, od))

						} else {
							return "", errors.New("order fmt error")
						}
					}
				case *JSON.Number_T:
					{
						if ch.Value > 0 {
							orders = append(orders, key+" asc")
						} else {
							orders = append(orders, key+" desc")
						}

					}
				case *JSON.Boolean_T:
					{
						if ch.Value {
							orders = append(orders, key+" asc")
						} else {
							orders = append(orders, key+" desc")
						}
					}
				}
			}
		}
	case JSON.Array:
		{
			for _, item := range vd {
				if i, y := item.(*JSON.String_T); !y {
					return "", errors.New("order fmt error")
				} else {
					orders = append(orders, i.Value)
				}
			}
		}
	case *JSON.Null_T:
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
