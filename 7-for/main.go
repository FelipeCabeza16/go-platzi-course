package main

import "fmt"

func main() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	for rango := range [3]int{} {
		fmt.Println(rango)
	}

	for _, valor := range [3]int{1, 2, 3} {
		fmt.Println(valor)
	}

	for indice, valor := range [3]int{1, 2, 3} {
		fmt.Println(indice, valor)
	}
}
