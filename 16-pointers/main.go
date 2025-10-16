package main

import "fmt"

func modificarArreglo(argumento *[5]int) {
	(*argumento)[0] = 42
}

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	modificarArreglo(&numeros)
	fmt.Println(numeros)
}
