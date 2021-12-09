package JSON

type Null struct{}

func (v Null) ToString() string {
	return "null"
}
