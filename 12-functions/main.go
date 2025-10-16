package main

import "fmt"

func main() {
	fmt.Println("Functions")

	suma := func(a int, b int) int {
		return a + b
	}

	fmt.Println(suma(1, 2))

	suma2 := func(a int, b int) int {
		return a + b
	}

	fmt.Println(suma2(1, 2))

	suma3 := func(a int, b int) int {
		return a + b
	}

	fmt.Println(suma3(1, 2))
}
