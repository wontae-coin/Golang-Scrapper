package main

import (
	"log"
	"os"
)

// * Go는 내장 타입으로 error라는 interface 타입을 갖는다
// 요 error 인터페이스로 에러 핸들링을 하는데, default 인터페이스는 아래와 같다
type error interface {
	Error() string
}

func main() {
	f, err := os.Open("./tmp.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())
}
