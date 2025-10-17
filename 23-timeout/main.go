package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timeout")

	timeout := 1 * time.Second
	channel := make(chan string)

	go func() {
		time.Sleep(timeout)
		channel <- "Mensaje"
	}()

	select {
	case mensaje := <-channel:
		fmt.Println(mensaje)
	case <-time.After(timeout):
		fmt.Println("Timeout")
	}

	fmt.Println("Fin del programa")
}
