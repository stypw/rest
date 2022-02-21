package gn

func (elt *element) IsValue() bool {
	if elt == nil {
		return false
	}
	switch elt.GetType() {
	case NullType:
		{
			return true
		}
	case BooleanType:
		{
			return true
		}
	case NumberType:
		{
			return true
		}
	case StringType:
		{
			return true
		}
	}
	return false
}
func (elt *element) RealValue() interface{} {
	if elt == nil {
		return nil
	}
	switch elt.GetType() {
	case NullType:
		{
			return elt.nullValue
		}
	case BooleanType:
		{
			return elt.boolValue
		}
	case NumberType:
		{
			return elt.numberValue
		}
	case StringType:
		{
			return elt.stringValue
		}
	}
	return nil
}
