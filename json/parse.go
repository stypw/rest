package JSON

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

func Parse(input string) (Object, error) {
	dec := json.NewDecoder(strings.NewReader(input))

	if t, err := dec.Token(); err == nil {
		if delim, yes := t.(json.Delim); yes {
			if string(delim) == "{" {
				var obj Object = make(Object)
				if err := obj.parse(dec); err == nil {
					return obj, nil
				}
			}
		}
	}
	return nil, errors.New("json string error")
}

func FromString(input string) (Object, error) {
	dec := json.NewDecoder(strings.NewReader(input))

	if t, err := dec.Token(); err == nil {
		if delim, yes := t.(json.Delim); yes {
			if string(delim) == "{" {
				var obj Object = make(Object)
				if err := obj.parse(dec); err == nil {
					return obj, nil
				}
			}
		}
	}
	return nil, errors.New("json string error")
}

func FromStream(stream io.ReadCloser) (Object, error) {
	dec := json.NewDecoder(stream)

	if t, err := dec.Token(); err == nil {
		if delim, yes := t.(json.Delim); yes {
			if string(delim) == "{" {
				var obj Object = make(Object)
				if err := obj.parse(dec); err == nil {
					return obj, nil
				}
			}
		}
	}
	return nil, errors.New("json string error")
}
