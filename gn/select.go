package gn

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(\w+)?((\[(\d+)\])|/|$)?`)

func (elt *element) TrySelect(path string) (Element, bool) {
	ret := re.FindAllStringSubmatch(path, -1)
	l := len(ret)
	var e Element = elt
	y := false
	for i := 0; i < l; i++ {
		keys := ret[i]
		kl := len(keys)
		if kl >= 2 && keys[1] != "" {
			if c, cy := e.TryGetProperty(keys[1]); cy {
				e = c
				y = true
			} else {
				y = false
				break
			}
		}
		if kl >= 5 && keys[4] != "" {
			if idx, ierror := strconv.Atoi(keys[4]); ierror == nil {
				if item, yes := e.TryGetElement(idx); yes {
					e = item
					y = true
				} else {
					y = false
					break
				}
			}
		}

	}
	return e, y
}
func (elt *element) Select(path string) Element {
	e, y := elt.TrySelect(path)
	if y {
		return e
	}
	return null
}
