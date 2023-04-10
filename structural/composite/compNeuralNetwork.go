package composite

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, other)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}
