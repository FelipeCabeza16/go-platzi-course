package main

import "fmt"

func main() {
	fmt.Println("Recursivity")

	resultado := factorial(5)
	fmt.Println(resultado)

	var fibonacci func(int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
	fmt.Println(fibonacci(7))

}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
