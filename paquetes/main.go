package main

import (
	"fmt"
	"os"
)

func main() {

	env := os.Getenv("HOME")
	if env == "" {
		fmt.Println("HOME is not set")
		return
	}

	fmt.Println("HOME is set to", env)
}
