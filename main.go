package main

import (
	"fmt"
	"os"
)

func main() {
	name := "John Doe"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greeting(name))
}

func greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
