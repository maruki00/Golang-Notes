package main

func Run(chan<- int) {

}

func main() {
	result := make(chan<- int, 0)


	select {
		case result
	}
}
