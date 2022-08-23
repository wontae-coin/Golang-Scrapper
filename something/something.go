package something

import "fmt"

// * 소문자로 시작하면 private
func sayBye() {
	fmt.Println("Bye")
}

// * 대문자로 시작하면 public
func SayHello() {
	fmt.Println("Hello")
}
