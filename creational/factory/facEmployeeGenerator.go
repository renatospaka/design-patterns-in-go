package factory

type EmployeeFacGen struct {
	Name, Position string
	AnnualIncome int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *EmployeeFacGen {
	return func(name string) *EmployeeFacGen {
		return &EmployeeFacGen{name, position, annualIncome}
	}
}

// struct
type EmployeeFacGenStruct struct {
	Position string
	AnnualIncome int
}

func (f *EmployeeFacGenStruct) Create(name string) *EmployeeFacGen {
	return &EmployeeFacGen{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFacGenStruct {
	return &EmployeeFacGenStruct{position, annualIncome}
}