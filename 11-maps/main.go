package main

import (
	"fmt"
	"maps"
)

func main() {
	mapa := make(map[string]int)
	mapa["Juan"] = 18
	mapa["Maria"] = 20
	fmt.Println(mapa)

	mapa2 := make(map[string]int)
	mapa2["amin"] = 4
	mapa2["Espinoza"] = 8
	fmt.Println(mapa2)

	if valor, existe := mapa2["amin"]; existe {
		fmt.Println("El valor de amin es", valor)
	} else {
		fmt.Println("El valor de amin no existe")
	}

	// delete(mapa, "amin") // Elimina únicamente la clave "amin".
	// clear(mapa) // Borra todos los registros.

	// Comparar mapas
	if maps.Equal(mapa, mapa2) {
		fmt.Println("mapa es igual a mapa2")
	} else {
		fmt.Println("mapa es diferente a mapa2")
	}
}
