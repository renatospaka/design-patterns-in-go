package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/composite"
)

func TestComposit() {
	fmt.Println("Composite => basic")
	drawing := composite.GraphicObject{Name: "My Drawing", Color: "", Children: nil}
	drawing.Children = append(drawing.Children, *composite.NewCircle("Red"))
	drawing.Children = append(drawing.Children, *composite.NewSquare("Yellow"))
	
	group := composite.GraphicObject{Name: "Group 1", Color: "", Children: nil}
	group.Children = append(group.Children, *composite.NewCircle("Blue"))
	group.Children = append(group.Children, *composite.NewSquare("Blue"))
	
	drawing.Children = append(drawing.Children, group)
	fmt.Println(drawing.String())

	fmt.Println("Composite => neural network")
	neuron1, neuron2 := &composite.Neuron{}, &composite.Neuron{}
	layer1, layer2 := composite.NewNeuronLayer(3), composite.NewNeuronLayer(4)
	composite.Connect(neuron1, neuron2)
	composite.Connect(neuron1, layer1)
	composite.Connect(layer2, neuron1)
	composite.Connect(layer1, layer2)
}
