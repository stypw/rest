package JSON

type Null_T struct{}

func Null() *Null_T {
	return &Null_T{}
}
func (v *Null_T) ToString() string {
	return "null"
}
