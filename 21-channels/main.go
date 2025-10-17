package main

import "fmt"

func main() {
	mensajes := make(chan string)
	go func() {
		mensajes <- "Ping"
	}()

	mensaje := <-mensajes
	fmt.Println(mensaje)
}
