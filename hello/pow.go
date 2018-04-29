package main

import "fmt"

func main() {
	var num int
	var pow int
	var result = 1

	num = 2
	pow = 4
	for i := 0; i < pow; i++ {
		result *= num
	}

	fmt.Printf("%d^%d = %d\n", num, pow, result)
}
