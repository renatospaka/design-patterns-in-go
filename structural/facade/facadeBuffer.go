package facade

type Buffer struct {
	width, height int
	buffer []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width * height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type ViewPort struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *ViewPort {
	return &ViewPort{buffer: buffer}
}

func (v *ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}
