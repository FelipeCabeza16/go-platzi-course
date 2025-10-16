package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const unicode = "áéíóú"

	fmt.Println("El unicode es:", unicode)
	fmt.Println("Len: ", len(unicode))

	for i := 0; i < len(unicode); i++ {
		fmt.Printf("%x ", unicode[i])
	}

	fmt.Println("Conteo de runas:", utf8.RuneCountInString(unicode))

	for idx, valorRuna := range unicode {
		fmt.Printf("%#U comienza en %d\n", valorRuna, idx)
	}
}
