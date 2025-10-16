package main

import "fmt"

func main() {
	numero := 9
	if numero < 0 {
		fmt.Println("Número negativo")
	} else if numero < 10 {
		fmt.Println("Número es de un solo dígito")
	} else {
		fmt.Println("Número mayor a nueve")
	}
}
