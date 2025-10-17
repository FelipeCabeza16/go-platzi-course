package main

import (
	"fmt"
	"time"
)

func main() {
	canalUno := make(chan string)
	canalDos := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		canalUno <- "Mensaje uno"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		canalDos <- "Mensaje dos"
	}()

	for i := 0; i <= 2; i++ {
		select {
		case mensajeUno := <-canalUno:
			fmt.Println(mensajeUno)
		case mensajeDos := <-canalDos:
			fmt.Println(mensajeDos)
		}
	}
}
