package JSON

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {

	const jsonStream = `
	{"Message": "Hello", "numbers":{"one":1},  "Null": null, "Number": 1.234}
`
	obj, err := Parse(jsonStream)
	if err != nil {
		fmt.Println(err)
	}
	if obj != nil {
		fmt.Println(obj.ToString())
	}

	var tst Value = String("123")

	fmt.Println(tst)

	s, _ := AsString(tst)
	fmt.Println(s)
}
