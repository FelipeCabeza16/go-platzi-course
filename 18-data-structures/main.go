package main

import "fmt"

type Kage string

type Ninja struct {
	Name string
	Age  int
	Kage Kage
}

const (
	Hokage     Kage = "Hokage"
	Kazekage   Kage = "Kazekage"
	Mizukage   Kage = "Mizukage"
	Raikage    Kage = "Raikage"
	Tsuchikage Kage = "Tsuchikage"
)

// tipos de datos personalizados
func CustomTypes() {
	Tsunade := Hokage
	fmt.Println("Tsunade es la: ", Tsunade)
	naruto := Ninja{
		Name: "Naruto",
		Age:  20,
		Kage: Hokage,
	}
	fmt.Println("Naruto es la: ", naruto)
}

func main() {
	CustomTypes()
}
