package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func dot(m, n mat.Matrix) mat.Matrix {
	r, _ := m.Dims()
	_, c := n.Dims()
	o := mat.NewDense(r, c, nil)
	o.Product(m, n)
	return o
}

func apply(fn func(i, j int, v float64) float64, m mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.Apply(fn, m)
	return o
}

func scale(s float64, m mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.Scale(s, m)
	return o
}

func multiply(m, n mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.MulElem(m, n)
	return o
}

func add(m, n mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.Add(m, n)
	return o
}

func subtract(m, n mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	o := mat.NewDense(r, c, nil)
	o.Sub(m, n)
	return o
}

func addScalar(scalar_value float64, m mat.Matrix) mat.Matrix {
	r, c := m.Dims()
	array := make([]float64, r*c)
	for x := 0; x < r*c; x++ {
		array[x] = scalar_value
	}
	n := mat.NewDense(r, c, array)
	return add(m, n)
}

func input(r int, c int) []float64 {
	var filled_array = make([]float64, r*c)
	for i := 0; i < 9; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanf("%f", &filled_array[i])
	}
	return filled_array
}

func main() {
	fmt.Println("Enter matrix elements(9): ")
	var input_array1 = input(3, 3)
	var input_array2 = input(3, 3)
	var m = mat.NewDense(3, 3, input_array1)
	var n = mat.NewDense(3, 3, input_array2)
	fmt.Println(mat.Formatted(m, mat.Prefix("    "), mat.FormatPython()))
	fmt.Println(mat.Formatted(n, mat.Prefix("    "), mat.FormatPython()))
	fmt.Printf(`Select Operation to perform:
	1.Addition
	2.Subtraction
	3.Dot Multiplication
	4.Standard Multiplication
	5.Scalar Addition
	6.Scalar Multiplication`)
}
