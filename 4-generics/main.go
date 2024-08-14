package main

type DefaultType interface {
	int | float32 | float64
}

func GetDefault[V DefaultType]() V {
	return 10
}

func main() {
	for 
}
