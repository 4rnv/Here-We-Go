package main
import ("fmt")

func octal() {
x := 15
y := 017
fmt.Println(x==y)
}

func slicemain() {
  slice_n := make([]int, 9, 12)
  fmt.Println(slice_n)
  fmt.Println(len(slice_n), cap(slice_n))
  slice_n[1],slice_n[4],slice_n[2],slice_n[8] = 9,3,2,5
  fmt.Println(slice_n)
  fmt.Println(len(slice_n), cap(slice_n))
  slice_n = append(slice_n,13,6,25,20)
  fmt.Println(slice_n)
  fmt.Println(len(slice_n), cap(slice_n))
}