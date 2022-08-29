package main

func main() {
	//* Anonymous function
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 3, 4, 5)
	println(result)
	
	//* first class function(일급함수)
	//* 다른 함수의 파라미터로 전달되거나, 리턴값으로 사용될 수 있다
	
	// 변수 add에 익명함수 전달
	add := func(i int, j int) int {
		return i + j
	}
	// 변수 r1에 add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)
	
	// 직접 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int {return x - y}, 10, 20)
	println(r2)
	
	
}

//* 변수 f의 타입을 func(int, int) int 라는 함수로 정한 것
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

//* type문을 사용한 함수 원형 정의
// type문은 struct, interface 등 custom type(user defined type)을 정의하기 위해 사용된다
// 함수 원형을 정의하는데에도 사용될 수 있다
//! 위 예제에서 func(x int, y int) int 함수 원형이 계속 반복되는데
type calculaor func(int, int) int
func calc2(f calculaor, a int, b int) int {
	result := f(a, b)
	return result
}