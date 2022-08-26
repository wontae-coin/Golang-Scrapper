package main

import "fmt"

// * struct 정의
type person struct {
	name string
	age  int
	//! 메소드를 정의하지 못하는데 어떻게 클래스를 구현하죠?
}

func main() {
	//* go의 struct(ts도 똑같았음)는 필드 데이터만을 가지며, 메소드를 가지지 않는다
	//* go에는 전통적인 oop언어가 갖는 클래스, 객체, 상속 개념이 없어서
	//* custom 타입을 정의하는 struct로 표현한다
	p := person{}

	p.name = "Lee"
	p.age = 10
	fmt.Println(p)
	//* constructor가 없는데도 초기화는 가능합니다
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p2)

}
