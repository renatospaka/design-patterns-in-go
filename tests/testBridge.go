package tests

import (
	"fmt"
	"time"

	"github.com/renatospaka/design-pattern-go/structural/bridge"
)

func TestBridge() {
	raster := bridge.RasterRender{}
	circle := bridge.NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
	fmt.Printf("BRIDGE => %s", "raster")


	vector := bridge.VectorRenderer{}
	circle2 := bridge.NewCircle(&vector, 5)
	circle2.Draw()
	circle2.Resize(2)
	circle2.Draw()
	fmt.Printf("BRIDGE => %s", "vector")

	time.Sleep(1 * time.Second)
}
