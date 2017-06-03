package calc

func Sub(a, b int, p chan int) {
	s := a - b
	p <- s
}
