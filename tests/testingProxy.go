package tests

import (
	"fmt"
	"log"

	"github.com/renatospaka/design-pattern-go/structural/proxy"
)

func TestProxy() {
	fmt.Println("Proxy => Protection Proxy")

	car := proxy.NewCarProxy(&proxy.Driver{12})
	car.Drive()

	car2 := proxy.NewCarProxy(&proxy.Driver{22})
	car2.Drive()

	fmt.Println("Proxy => Virtual Proxy")
	bmp := proxy.NewBitmap("demo.png")
	DrawImage(bmp)

	lbmp := proxy.NewLazyBitmap("demo.png")
	DrawImage(lbmp)
	DrawImage(lbmp)
}

func DrawImage(image proxy.Image) {
	log.Println("About to draw the image")
	image.Draw()
	log.Println("Done drawing the image\n")
}
