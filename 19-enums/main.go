package main

import "fmt"

type serverState int

const (
	idle serverState = iota
	connected
	error
	retry
)

var stateName = map[serverState]string{
	idle:      "idle",
	connected: "connected",
	error:     "error",
	retry:     "retry",
}

func printState(state serverState) string {
	switch state {
	case idle:
		return "el servidor está inactivo"
	case connected:
		return "el servidor está conectado"
	case error:
		return "el servidor está en error"
	case retry:
		return "el servidor está en retry"
	}
	panic("estado desconocido")
}

func main() {
	fmt.Println("Enums")

	redServer := printState(connected)
	fmt.Println(redServer)

	redServer = printState(error)
	fmt.Println(redServer)

	redServer = printState(retry)
	fmt.Println(redServer)
}
