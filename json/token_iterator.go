package JSON

import (
	"encoding/json"
	"io"
)

type _BooleanHandle func(bool) bool
type _StringHandle func(string) bool
type _NumberHandle func(float64) bool
type _NullHandle func() bool
type _ObjectStartHandle func() bool
type _ObjectEndHandle func() bool
type _ArrayStartHandle func() bool
type _ArrayEndHandle func() bool

type tokenIterator struct {
	dec               *json.Decoder
	booleanHandle     _BooleanHandle
	stringHandle      _StringHandle
	numberHandle      _NumberHandle
	nullHandle        _NullHandle
	objectStartHandle _ObjectStartHandle
	objectEndHandle   _ObjectEndHandle
	arrayStartHandle  _ArrayStartHandle
	arrayEndHandle    _ArrayEndHandle
}

func (interator *tokenIterator) run() {
	for {
		t, err := interator.dec.Token()
		switch tv := t.(type) {
		case json.Delim:
			{

				st := string(tv)
				switch st {
				case "{":
					{
						if interator.objectStartHandle != nil {
							if !(interator.objectStartHandle()) {
								return
							}
						}
					}
				case "}":
					{
						if interator.objectEndHandle != nil {
							if !(interator.objectEndHandle()) {
								return
							}
						}
					}
				case "[":
					{
						if interator.arrayStartHandle != nil {
							if !(interator.arrayStartHandle()) {
								return
							}
						}
					}
				case "]":
					{
						if interator.arrayEndHandle != nil {
							if !(interator.arrayEndHandle()) {
								return
							}
						}
					}
				}
			}
		case string:
			{
				if interator.stringHandle != nil {
					if !(interator.stringHandle(tv)) {
						return
					}
				}
			}
		case bool:
			{
				if interator.booleanHandle != nil {
					if !(interator.booleanHandle(tv)) {
						return
					}
				}
			}
		case float64:
			{
				if interator.numberHandle != nil {
					if !(interator.numberHandle(tv)) {
						return
					}
				}
			}
		case nil:
			{
				if interator.nullHandle != nil {
					if !(interator.nullHandle()) {
						return
					}
				}
			}
		}
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
	}
}
