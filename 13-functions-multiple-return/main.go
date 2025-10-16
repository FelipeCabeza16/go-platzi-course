package main

import "fmt"

func main() {
	fmt.Println("Functions Multiple Return")

	suma, resta := sumaYResta(1, 2)
	fmt.Println(suma, resta)
}

func sumaYResta(a int, b int) (int, int) {
	return a + b, a - b
}
