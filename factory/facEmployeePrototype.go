package factory

type EmployeeFacProto struct {
	Name, Position string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
)

func NewEmployeeFacProto(role int) *EmployeeFacProto {
	switch role {
	case Developer:
		return &EmployeeFacProto{"", "developer", 60000}
	case Manager:
		return &EmployeeFacProto{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}