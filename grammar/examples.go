package main

import (
	"fmt"
)

func octal() {
	x := 15
	y := 017
	fmt.Println(x == y)
}

func slicemain() {
	slice_n := make([]int, 9, 12)
	fmt.Println(slice_n)
	fmt.Println(len(slice_n), cap(slice_n))
	slice_n[1], slice_n[4], slice_n[2], slice_n[8] = 9, 3, 2, 5
	fmt.Println(slice_n)
	fmt.Println(len(slice_n), cap(slice_n))
	slice_n = append(slice_n, 13, 6, 25, 20)
	fmt.Println(slice_n)
	fmt.Println(len(slice_n), cap(slice_n))
}

func pattern() {
	var n, i, j int
	fmt.Println("Enter le number")
	fmt.Scanf("%d", &n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
	for i = n - 1; i > 0; i-- {
		for j = i; j > 0; j-- {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func prime() {
	var n, i int
	c := 0
	fmt.Println("Enter le number")
	fmt.Scanf("%d", &n)
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			c += 1
		}
	}
	if c == 2 {
		fmt.Printf("Prime Number")
	} else {
		fmt.Printf("Not a Prime Number, has %d factors", c)
	}
}

func Factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

func FactorialMain() {
	var n int
	fmt.Println("Enter le number")
	fmt.Scanf("%d", &n)
	x := Factorial(n)
	fmt.Printf("%d", x)
}
