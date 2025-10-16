package main

import "fmt"

func main() {
	fmt.Println("Variadic Functions")

	suma := func(numeros ...int) int {
		total := 0
		for _, numero := range numeros {
			total += numero
		}
		return total
	}

	fmt.Println(suma(1, 2, 3, 4, 5))
}
