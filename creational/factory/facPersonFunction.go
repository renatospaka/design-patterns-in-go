package factory

type PersonFuncFactory struct {
	Name string
	Age int
	EyeCount int
}

func NewPersonFuncFactory(name string, age int) *PersonFuncFactory {
	if age < 16 {
		// ...
	}
	
	return &PersonFuncFactory{name, age, 2}
}
