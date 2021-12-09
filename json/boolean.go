package JSON

type Boolean bool

func (v Boolean) ToString() string {
	if v {
		return "true"
	} else {
		return "false"
	}
}

func GetBoolean(v Value) bool {
	if b, y := v.(Boolean); y {
		return bool(b)
	}
	return false
}

func TryGetBoolean(v Value) (bool, bool) {
	if b, y := v.(Boolean); y {
		return bool(b), true
	}
	return false, false
}
