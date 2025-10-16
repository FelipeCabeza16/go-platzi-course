package main

import (
	"fmt"
	"time"
)

func main() {
	switch numero := 9; numero {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	}

	var I int = 2
	switch I {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}
	// TIEMPO

	tiempo := time.Now()
	switch {
	case tiempo.Hour() < 12:
		fmt.Println("Debes decir buenos días")
	default:
		fmt.Println("Debes decir buenas tardes")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("A descansar")
	default:
		fmt.Println("Toca grabar más cursos en Platzi")
	}
}
