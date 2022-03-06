package kv

import (
	"strconv"
	"strings"
)

func TrySelect(elt Element, path string) (Element, bool) {
	var e Element = elt
	keys := strings.Split(path, "/")
	for _, k := range keys {
		if k == "" {
			continue
		}
		indexes := strings.Split(k, "[")
		first := indexes[0]

		if c, cy := e.TryGetProperty(first); cy {
			e = c
		} else {
			return elt, false
		}

		if len(indexes) < 2 {
			continue
		}
		for _, str := range indexes[1:] {
			if str == "" {
				continue
			}
			l := len(str)
			if l < 2 {
				return elt, false
			}
			if str[l-1] != ']' {
				return elt, false
			}
			idx, ierror := strconv.Atoi(strings.TrimRight(str, "]"))
			if ierror != nil {
				return elt, false
			}
			item, yes := e.TryGetElement(idx)
			if !yes {
				return elt, false
			}
			e = item
		}
	}

	return e, true
}
func Select(elt Element, path string) Element {
	e, y := elt.TrySelect(path)
	if y {
		return e
	}
	return undefined
}
