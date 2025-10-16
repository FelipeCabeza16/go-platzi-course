package arreglos

import "fmt"

func main() {
	var arreglo [3]int
	arreglo[0] = 1
	arreglo[1] = 2
	arreglo[2] = 3
	fmt.Println(arreglo)

	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)

	arreglo3 := [...]int{1, 2, 3}
	fmt.Println(arreglo3)

	arreglo4 := [3]int{1: 2, 2: 3}
	fmt.Println(arreglo4)

	arreglo5 := [3]int{1, 2, 3}
	fmt.Println(arreglo5)

	arreglo6 := [3]int{1, 2, 3}
	fmt.Println(arreglo6)

	arreglo7 := [3]int{1, 2, 3}
	fmt.Println(arreglo7)
}
