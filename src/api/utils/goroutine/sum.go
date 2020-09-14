package goroutine

func sum(array []int, c chan int) {
	sum := 0

	for _, val := range array {
		sum += val
	}

	c <- sum
}

func returnSum() (int, int) {
	var array = []int{10, 48, 2, 4, 1}

	c := make(chan int)

	go sum(array[len(array)/2:], c)
	go sum(array[:len(array)/2], c)

	x, y := <-c, <-c
	return x, y
}
