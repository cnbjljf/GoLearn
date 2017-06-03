package calc

func Add(a, b int, p chan int) {
	s := a + b
	p <- s
}
