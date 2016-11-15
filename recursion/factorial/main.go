package main

import "fmt"

func main() {
	fmt.Printf("The factorial of 3 is: %d\n", fact(3))
	fmt.Printf("The factorial of 4 is: %d\n", fact(4))
	fmt.Printf("The factorial of 5 is: %d\n", fact(5))
	fmt.Printf("The factorial of 10 is: %d", fact(10))
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return fact(n - 1) * n
	}
}
