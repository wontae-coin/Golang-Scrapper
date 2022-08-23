package main

//* main entry point for the compiled code
//* If not using the compiled code and wanting to share the code in any form like library etc.
// not necessary to name this 'main.go'

import (
	"fmt"
	"strings"

	"github.com/wontae/learngo/something"
)

// * GO func can have many return values
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) string {
}

func main() {
	fmt.Println("Hi")
	something.SayHello()
	const name = "wontae"
	println(multiply(2, 2))
	length, _ := lenAndUpper("wontae")
	fmt.Println(length)

	repeatMe("wontea", "lynn", "dal")
}
