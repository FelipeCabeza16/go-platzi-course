package main

import "fmt"

func main() {
	fmt.Println("Valores")

	var a int = 1
	var b int = 2
	var c int = a + b
	fmt.Println(c)

	var entero = 10
	var doble = 2.5

	resultado := entero + int(doble)
	fmt.Println(resultado)

	var nombre string = "Juan"
	apellido := "Perez"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)
}
