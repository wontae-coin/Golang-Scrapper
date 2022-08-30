package main

import "fmt"

// * 타 언어의 OOP가 필드와 메서드를 함께 갖는 것과 달리 Go는 struct가 필드를 가지며, 메소드는 별로도 분리되어 정의한다
type Rect struct {
	width, height int
}

// * 메소드를 정의할 때 그 메소드가 어떤 struct를 위한 메소드인지 정의한다
// * Rect의 area() 메소드
// * func (receiver Struct) method() return {}
func (r Rect) area() int {
	return r.width * r.height
}

func GetMethod() {
	rect := Rect{10, 20}
	// 메소드가 선언된 이후에는 Rect struct의 객체는 자유롭게 메소드를 사용할 수 있다
	area := rect.area()
	fmt.Println(area)

}
