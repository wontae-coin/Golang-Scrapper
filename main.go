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

// * Naked function function
func lenAndLower(name string) (length int, lowercase string) {
	// These variables below are declared as arguments
	// Not declaring, but modifying
	length = len(name)
	lowercase = strings.ToLower(name)
	// Naked  return
	return
}

// * DEFER
// 함수가 실행 종료되었을 때 추가적인 행동을 명령할 수 있다
func lenAndLowerDefer(name string) (int, string) {
	// Decorator처럼 쓸 수도, 콜백처럼 쓸 수도 있다. 캐시를 삭제하다던가 사용한 파일을 삭제한다던가
	defer fmt.Println("This function is done")
	return lenAndLower(name)
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	for number := range numbers {
		fmt.Println(number)
	}
	return 1
}

func main() {
	fmt.Println("Hi")
	something.SayHello()
	const name = "wontae"
	println(multiply(2, 2))
	length, _ := lenAndUpper("wontae")
	fmt.Println(length)
	lenAndLowerDefer("hongjun")
	repeatMe("wontea", "lynn", "dal")
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
