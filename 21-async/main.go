package main

import (
	"fmt"
	"time"
)

func imprimirHola(str string) {
	fmt.Println(str)
}

func main() {
	imprimirHola("modo directo")

	go imprimirHola("modo asíncrono")
	// si comento esto, el programa termina antes de que se imprima el mensaje asíncrono
	time.Sleep(1 * time.Second)

	fmt.Println("Fin del programa")
}
