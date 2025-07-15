package main

import (
	"fmt"
	"os"

	"github.com/samber/lo"
)

const (
	defaultName     = "John Doe"
	greetingPattern = "Hello, %s!"
)

func main() {
	fmt.Println(greeting(name(os.Args)))
}

func greeting(name string) string {
	return fmt.Sprintf(greetingPattern, name)
}

func name(args []string) string {
	return lo.IfF(len(args) > 1, func() string { return args[1] }).Else(defaultName)
}
