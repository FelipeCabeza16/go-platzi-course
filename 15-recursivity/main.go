package main

import "fmt"

func main() {
	fmt.Println("Recursivity")

	resultado := factorial(5)
	fmt.Println(resultado)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
