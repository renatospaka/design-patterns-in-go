package builder

type PersonF struct {
	name, position string
}

type personMod func(*PersonF)
type PersonBuilderF struct {
	actions []personMod
}

func (b *PersonBuilderF) Called(name string) *PersonBuilderF {
	b.actions = append(b.actions, func(p *PersonF) {
		p.name = name
	})
	return b
}

func (b *PersonBuilderF) WorksAsA(position string) *PersonBuilderF {
	b.actions = append(b.actions, func(p *PersonF) {
		p.position = position
	})
	return b
}

func (b *PersonBuilderF) Build() *PersonF {
	p := PersonF{}
	for _, i := range b.actions {
		i(&p)
	}

	return &p
}