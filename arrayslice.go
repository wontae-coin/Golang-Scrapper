package main

import "fmt"

func main() {
	//* go array의 length는 최대값이다. 작아도 에러는 발생하지 않는다. 넘으면 발생
	names := [5]string{"a", "b", "c"}
	fmt.Println(names)
	//* length를 제한하기를 원하지 않는다면 slice(= python의 리스트)
	// length만 빼주면 된다
	calls := []string{"x", "y", "z"}
	// append falls into unmodified state, in other words, it returns the new state
	calls = append(calls, "flynn")
	fmt.Println(calls)
}
