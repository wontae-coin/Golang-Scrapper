package main

import "fmt"

// * struct 정의
type person struct {
	name string
	age  int
	//! 메소드를 정의하지 못하는데 어떻게 클래스를 구현하죠?
}
//* go의 struct(ts도 똑같았음)는 필드 데이터만을 가지며, 메소드를 가지지 않는다
//* go에는 전통적인 oop언어가 갖는 클래스, 객체, 상속 개념이 없어서
//* custom 타입을 정의하는 struct로 표현한다

func GetStruct() {
	//* 객체 할당 먼저, 이후에 필드값 채우기
	p1 := person{}
	p1.name = "Lee"
	p1.age = 10
	fmt.Println(p1)

	//* 필드값 채우는 두번째 방법
	var p2 person
	p2 = person{"Bob", 20}
	fmt.Println(p2)

	//* 할당할 때 필드명을 지정하는 방법
	// 일부 필드가 생략되었다면 zero value로 들어간다
	p3 := person{name: "Sean", age: 50}
	fmt.Println(p3)

	//* new 내장함수를 사용하는 법
	// 모든 필드를 zero value로 초기화하고 person객체의 포인터 * 를 리턴한다
	p4 := new(person)
	// 객체가 포인터인 경우에도 필드 액세스 시 .를 사용하는데, 자동으로 dereferencing 된다
	p4.name = "Lee"

	//* 생성자 호출
	dic := newDict()
	dic.data[1] = "A"
}

// * 생성자 constructor
// 선언될 때의 자동으로 실행되는 default 동작
type dict struct {
	data map[int]string
}

// * 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}