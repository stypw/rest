package cf

import (
	"fmt"
	"rest/gn"
	"sync"
)

var lock sync.Mutex
var cf gn.Element = gn.Null

func GetConfig() gn.Element {
	if cf == gn.Null {
		lock.Lock()
		if cf == gn.Null {
			cf = gn.FromFile("./config.json")
		}
		if cf == gn.Null {
			cf = gn.FromFile("./app.json")
		}
		if cf == gn.Null {
			cf = gn.FromFile("./application.json")
		}
		if cf == gn.Null {
			cf = gn.FromFile("./cf.json")
		}
		if cf == gn.Null {
			fmt.Println("could not find configfile:[config.josn | app.json | application.json | cf.json]")
		}
		lock.Unlock()
	}
	return cf
}
