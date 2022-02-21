package gn

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func parseObject(dc *json.Decoder) Element {
	obj := NewObject()
	k := ""
	for {
		t, err := dc.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		ct := false
		switch tv := t.(type) {
		case json.Delim:
			{
				switch string(tv) {
				case "{":
					{
						if k != "" {
							obj.Set(k, parseObject(dc))
							k = ""
							ct = true
						}
					}
				case "[":
					{
						if k != "" {
							obj.Set(k, parseArray(dc))
							k = ""
							ct = true
						}
					}
				}
			}
		case string:
			{
				if k != "" {
					obj.Set(k, NewString(tv))
					k = ""
				} else {
					k = tv
				}
				ct = true
			}
		case float64:
			{
				if k != "" {
					obj.Set(k, NewNumber(tv))
					k = ""
					ct = true
				}
			}
		case bool:
			{
				if k != "" {
					obj.Set(k, NewBoolean(tv))
					k = ""
					ct = true
				}
			}
		case nil:
			{
				if k != "" {
					obj.Set(k, NewNull())
					k = ""
					ct = true
				}
			}
		}
		if !ct {
			break
		}
	}
	return obj
}
func parseArray(dc *json.Decoder) Element {
	arr := NewArray()
	for {
		t, err := dc.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		ct := true
		switch tv := t.(type) {
		case json.Delim:
			{
				switch string(tv) {
				case "{":
					{
						arr.Push(parseObject(dc))
					}
				case "[":
					{
						arr.Push(parseArray(dc))
					}
				default:
					{
						ct = false
					}
				}
			}
		case string:
			{
				arr.Push(NewString(tv))
			}
		case float64:
			{
				arr.Push(NewNumber(tv))
			}
		case bool:
			{
				arr.Push(NewBoolean(tv))
			}
		case nil:
			{
				arr.Push(NewNull())
			}
		default:
			{
				ct = false
			}
		}
		if !ct {
			break
		}
	}
	return arr
}

func parse(dc *json.Decoder) Element {

	t, err := dc.Token()
	if err == io.EOF {
		return null
	}
	if err != nil {
		return null
	}
	switch tv := t.(type) {
	case json.Delim:
		{
			switch string(tv) {
			case "{":
				{
					return parseObject(dc)
				}
			case "[":
				{
					return parseArray(dc)
				}
			}
		}
	}
	return null
}

func FromStream(stream io.Reader) Element {
	dc := json.NewDecoder(stream)
	return parse(dc)
}

func FromString(input string) Element {
	return FromStream(strings.NewReader(input))
}

func FromFile(path string) Element {
	ct, er := ioutil.ReadFile("./config.json")
	if er != nil {
		return null
	}
	return FromString(string(ct))
}
