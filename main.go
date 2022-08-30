package main

//* main entry point for the compiled code
//* If not using the compiled code and wanting to share the code in any form like library etc.
//* then is is not necessary to name this 'main.go'

import (
	"fmt"
)

func main() {
	result := LearnMap()
	fmt.Println(result)
}
