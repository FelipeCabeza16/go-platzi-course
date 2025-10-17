package main

import (
	"fmt"
	"time"
)

func worker(id int, tareas <-chan int, resultados chan<- int) {
	for tarea := range tareas {
		fmt.Printf("Worker %d started task %d\n", id, tarea)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished task %d\n", id, tarea)
		resultados <- tarea * 2
	}
}

func main() {
	fmt.Println("Worker")

	tareas := make(chan int, 100)
	resultados := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go worker(i, tareas, resultados)
	}

	for i := 0; i < 10; i++ {
		tareas <- i
	}

	close(tareas)

	for i := 0; i < 10; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	close(resultados)
}
