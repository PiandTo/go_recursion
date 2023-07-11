package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int64) int64 {
	var sum int64
	if n == 1 {
		sum += 1
	} else if n > 1 {
		sum += (fibonacci(n-2) + fibonacci(n-1))
	}
	return sum
}

func main() {
	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}
