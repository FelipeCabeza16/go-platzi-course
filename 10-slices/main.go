package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	var arregloCadenas []string
	fmt.Println("datos", "contenido", arregloCadenas, "condicion ", arregloCadenas == nil, "longitud", len(arregloCadenas), "capacidad", cap(arregloCadenas))

	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "Hola"
	arregloCadenas[1] = "Mundo"
	arregloCadenas[2] = "Platzi"
	fmt.Println("datos", "contenido", arregloCadenas, "condicion ", arregloCadenas == nil, "longitud", len(arregloCadenas), "capacidad", cap(arregloCadenas))

	arregloCadenas = append(arregloCadenas, "Platzi")
	fmt.Println("datos", "contenido", arregloCadenas, "condicion ", arregloCadenas == nil, "longitud", len(arregloCadenas), "capacidad", cap(arregloCadenas))

	segundoSlice := arregloCadenas[1:3]
	fmt.Println("datos", "contenido", segundoSlice, "condicion ", segundoSlice == nil, "longitud", len(segundoSlice), "capacidad", cap(segundoSlice))
	if slices.Equal(segundoSlice, []string{"Mundo", "Platzi"}) {
		fmt.Println("segundoSlice es igual a []string{\"Mundo\", \"Platzi\"}")
	}

	if slices.Contains(segundoSlice, "Platzi") {
		fmt.Println("segundoSlice contiene Platzi")
	}

	if slices.Contains(segundoSlice, "Python") {
		fmt.Println("segundoSlice contiene Python")
	}
}
