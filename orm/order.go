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
				case JSON.String:
					{
						var od = strings.ToLower(strings.Trim(string(ch), ""))
						if od == "asc" || od == "desc" {
							orders = append(orders, fmt.Sprintf("%s %s", key, od))
							orders = append(orders, fmt.Sprintf("%s %s", key, od))

						} else {
							return "", errors.New("order fmt error")
						}
					}
				case JSON.Number:
					{
						if ch > 0 {
							orders = append(orders, key+" asc")
						} else {
							orders = append(orders, key+" desc")
						}

					}
				case JSON.Boolean:
					{
						if ch {
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
				if i, y := item.(JSON.String); !y {
					return "", errors.New("order fmt error")
				} else {
					orders = append(orders, string(i))
				}
			}
		}
	case JSON.Null:
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
