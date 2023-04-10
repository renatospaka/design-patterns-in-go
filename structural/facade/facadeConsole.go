package facade

type Console struct {
	buffer []*Buffer
	viewports []*ViewPort
	offset int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{
		[]*Buffer{b},
		[]*ViewPort{v},
		0,
	}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}
