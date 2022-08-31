package main

import (
	"fmt"
	"math"
)

// * struct가 필드들의 집합체라면, interface는 메소드들의 집합체
// * type이 구현해야 하는 메소드의 원형(prototype)dmf wjddmlgksek
// * 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 그 인터페이스가 갖는 모든 메소드들을 구현하면 된다
type Shape interface {
	area() float64
	permieter() float64
}

// * Rect 정의
type Rect struct {
	width, height float64
}

// * Circle 정의
type Circle struct {
	radius float64
}

//* Rect 타입에 대한 Shape 인터페이스 구현

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) permieter() float64 {
	return 2 * (r.width + r.height)
}

//* Circle type에 대한 Shape 인터페이스 구현

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) permieter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)

	//* Interface type (empty interface, interface{})
	// 메서드를 전혀 갖지 않는 인터페이스인데, Go의 모든 타입은 적어도 0개의 메서드를 구현하므로
	// Go의 모든 타입을 나타내기 위해 사용되는, 즉 어떠한 타입도 담을 수 있는 다이나믹 타입입니다
	var x interface{}
	x = 1
	x = "Tom"

	//* Type assertion
	y := x.(string)
	// x는 정수를 담았다가 문자열을 담는데, 아무런 문제없이 문자열 Tom을 출력한다
	printIt(x)
	println(y)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func printIt(v interface{}) {
	fmt.Println(v)
}

//* Type assertion
