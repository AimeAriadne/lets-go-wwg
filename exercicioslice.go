package main

import (
	"fmt"
)

// declare duas slices de int e preencha 6 posições de cada uma.
// junte as slices formando uma terceira slide.
// print as três slices

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1)
	slice2 := []int{7, 8, 9, 10, 11, 12}
	fmt.Println(slice2)

	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)
}
