package main

import (
	"fmt"
	"math/rand"
)

type Layer struct {
	weights []float64
	bias    float64
}

func relu(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}

func newLayer(inputSize int) Layer {
	weights := make([]float64, inputSize)
	for i := range weights {
		weights[i] = rand.Float64()
	}
	bias := rand.Float64()
	return Layer{weights: weights, bias: bias}
}

func layer1(x int, y int) float64 {
	config_param := rand.Float64()
	return config_param * float64(x+y)
}

func layer2(x float64, n int) float64 {
	config_param := rand.Float64()
	return config_param * (x * float64(n))
}

func (l Layer) forward(inputs []float64) float64 {
	sum := l.bias
	for i, weight := range l.weights {
		sum += weight * inputs[i]
	}
	return relu(sum)
}

func main() {
	layer1 := newLayer(2)
	layer2 := newLayer(1)
	k := 100.0
	y := 20.0
	inputs := []float64{k, y}
	output1 := layer1.forward(inputs)
	fmt.Println("Output from Layer 1:", output1)

	output2 := layer2.forward([]float64{output1})
	fmt.Println("Output from Layer 2:", output2)
}
