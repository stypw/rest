package kv

const (
	UndefinedType = iota
	NullType
	BooleanType
	NumberType
	StringType
	ObjectType
	ArrayType
)

//注意事项，无效的Element不能通过 Element==nil来判断，应该调用Element.IsValid()来进行判断
type Element interface {
	GetType() int

	//null才能正确返回，都返回nil
	GetNull() interface{}
	TryGetNull() (interface{}, bool)

	//boolean才能正确返回,其它返回false
	GetBoolean() bool
	TryGetBoolean() (bool, bool)

	//number才能正确返回，其它返回0
	GetNumber() float64
	TryGetNumber() (float64, bool)

	//string才能正确返回，其它返回空字符串""
	GetString() string
	TryGetString() (string, bool)

	//返回object[k],object类型且k匹配才能正确返回，其它返回undefined
	GetProperty(k string) Element
	TryGetProperty(k string) (Element, bool)

	//返回array[idx],array类型且idx在索引范围内才能正确返回，其它返回undefined
	GetElement(idx int) Element
	TryGetElement(idx int) (Element, bool)

	//object才能调用，返回调用对象以进行链式调用
	Set(k string, v Element) Element

	//array才能调用，返回调用对象以进行链式调用
	Push(v Element) Element

	//路径错误也会返回undefined，不能通过 == nil判断，请调用IsValid();
	//示例：elt.Select("/info/skills[0][1]/golang") = elt.GetProperty("info").GetProperty("skills").GetElement(0).GetElement(1).GetProperty("golang")
	Select(path string) Element
	TrySelect(path string) (Element, bool)

	//array类型才会返回[]Element，其它返回nil
	ArrayEnumerator() []Element

	//object才会返回map[string]Element，其它返回nil
	ObjectEnumerator() map[string]Element

	//返回Json字符串，对于StringType类型，返回的字符串两边会加上引号""
	ToString() string

	//null,boolean,number,string返回具体值,其它返回nil
	GetValue() interface{}

	//null,boolean,number,string返回true
	IsValue() bool

	//null,boolean,number,string,object,array返回true
	IsValid() bool
}
