package main

import "fmt"

type DefaultType interface {
	int | float32 | float64
}

func GetDefault[V DefaultType](items []V) V {
	var result V
	for _, x := range items {
		result += x
	}

	return result
}

func main() {
	items1 := []int{1, 23, 54, 6456, 45645, 4564}
	items2 := []float32{1.345, 23.345, 54.234, 6456.78, 45645.78, 4564.00}
	fmt.Println("result : ", GetDefault[int](items1))
	fmt.Println("result : ", GetDefault[float32](items2))

}
