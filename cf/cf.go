package cf

import (
	"fmt"
	"rest/kv"
	"sync"
)

var lock sync.Mutex
var cf kv.Element = kv.Null

func GetConfig() kv.Element {
	if cf == kv.Null {
		lock.Lock()
		if cf == kv.Null {
			cf = kv.FromFile("./config.json")
		}
		if cf == kv.Null {
			cf = kv.FromFile("./app.json")
		}
		if cf == kv.Null {
			cf = kv.FromFile("./application.json")
		}
		if cf == kv.Null {
			cf = kv.FromFile("./cf.json")
		}
		if cf == kv.Null {
			fmt.Println("could not find configfile:[config.josn | app.json | application.json | cf.json]")
		}
		lock.Unlock()
	}
	return cf
}
