package prototype

import (
	"bytes"
	"encoding/gob"
)

// Ready to go prototypes for employees
var mainOffice = Employee{"", Address{0, "av. Paulista, 1344", "São Paulo"}}
var auxOffice = Employee{"", Address{0, "av. Tamboré, 543", "Tamboré"}}

type Address struct {
	Suit int
	StreetAddress, City string
}

type Employee struct {
	Name string
	Office Address
}

func newEmployee(proto *Employee, name string, suit int) *Employee {
	newEmp := proto.DeepCopy()
	newEmp.Name = name
	newEmp.Office.Suit = suit
	return newEmp
}

func NewMainOfficeEmployee(name string, suit int) *Employee {
	return newEmployee(&mainOffice, name, suit)
}

func NewAuxOfficeEmployee(name string, suit int) *Employee {
	return newEmployee(&auxOffice, name, suit)
}

func (p* Employee) DeepCopy() *Employee {
	// no error handling in this sample
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(p)

	decoder := gob.NewDecoder(&buffer)
	newEmp := Employee{}
	_ = decoder.Decode(&newEmp)

	return &newEmp
}
