package rt

type RouterFactory interface {
	Create() Router
}

type FactoryFunc func() Router

func (f FactoryFunc) Create() Router {
	return f()
}
func NewFactory(f FactoryFunc) RouterFactory {
	return FactoryFunc(f)
}
