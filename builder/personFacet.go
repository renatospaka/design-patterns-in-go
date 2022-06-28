package builder

type PersonFa struct {
// addresses info
StreetAddress, Postcode, City string

// jon info
CompanyName, Position string
AnnualIncome int
}

type PersonBuilderFa struct {
	person *PersonFa
}

func NewPersonBuilderFa() *PersonBuilderFa {
	return &PersonBuilderFa{
		person: &PersonFa{},
	}
}

type PersonAddressBuilder struct {
	PersonBuilderFa
}

type PersonJobBuilder struct {
	PersonBuilderFa
}

func (b *PersonBuilderFa) Lives() *PersonAddressBuilder {
	// Due to PersonAddressBuilder agregates PersonBuilderF
	// {person} pointer has to be initialized here
	return &PersonAddressBuilder{*b}
}

func (p *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	p.person.StreetAddress = streetAddress
	return p
}

func (p *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	p.person.City = city
	return p
}

func (p *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	p.person.Postcode = postcode
	return p
}

func (b *PersonBuilderFa) Works() *PersonJobBuilder {
	// Due to PersonJobBuilder agregates PersonBuilderF
	// {person} pointer has to be initialized here
	return &PersonJobBuilder{*b}
}

func (p *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	p.person.CompanyName = companyName
	return p
}

func (p *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	p.person.Position = position
	return p
}

func (p *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	p.person.AnnualIncome = annualIncome
	return p
}

func (b *PersonBuilderFa) Build() *PersonFa {
	return b.person
}