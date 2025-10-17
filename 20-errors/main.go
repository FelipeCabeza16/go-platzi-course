package main

import (
	"fmt"
)

var errorDeCafe = fmt.Errorf("Ya no hay café")
var errorDeEnergia = fmt.Errorf("Ya no hay electricidad")

func hacerCafe(args int) error {
	if args == 2 {
		return errorDeCafe
	} else if args == 3 {
		return errorDeEnergia
	}
	return nil
}

func main() {
	for i := 0; i < 5; i++ {
		err := hacerCafe(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Ya hay café")
		}
	}
}
