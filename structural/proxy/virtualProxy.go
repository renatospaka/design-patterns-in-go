package proxy

import "log"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	log.Println("Drawing the image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	log.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

type LazyBitmap struct {
	filename string
	bitmap *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap((l.filename))
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}
