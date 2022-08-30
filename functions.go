package main

import (
	"fmt"
	"strings"
)

// * 함수
func multiply(a int, b int) int {
	return a * b
}

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

//! ---------------------------------------------

// * Pass by value
// ! boolean, numeric(int, float), string
func say(msg string) {
	println(msg)
}

// * Pass by reference
// ! function, pointer, slice, array
func bye(msg *string) {
	// 파라미터가 포인터임을 표시하는 법 *
	println("받아온건 주소!", msg)
	println("주소가 가졌던/가지는 값을 가져오려면 dereferencing:", *msg)
	*msg = "Changed"
}

// * Variadic Function (가변인자함수)
// * arguments들을 모아 array로 만들어 줍니다
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println("Hi")
	const name = "wontae"
	println(multiply(2, 2))
	length, _ := lenAndUpper("wontae")
	fmt.Println(length)
	lenAndLowerDefer("hongjun")

	//! ---------------------------------------------
	repeatMe("wontea", "lynn", "dal")

	msg := "Hello"
	say(msg)
	//* pointer문법, RUST의 reference(빌림)과 다르다
	//* msg의 값을 가져오는게 아니라 msg 변수의 주소를 가져온다
	bye(&msg)
	println(msg)
}
