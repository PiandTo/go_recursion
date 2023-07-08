package main

import (
	"fmt"
)

func factorial(n int64) int64 {
	sum := int64(1)
	if n > 1 {
		sum = n * factorial(n-1)
	}
	return sum
}

func main() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}