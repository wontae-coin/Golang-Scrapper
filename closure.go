package main
//* Go에서 함수는 Closure처럼 사용될 수 있다
// Closure: 함수의 기본 렉시컬 환경에서 참조할 변수를 찾지 못하면 상위 렉시컬 환경을 참조하는 것

//* int를 리턴하는 익명함수 func() int를 리턴한다
// Go에서도 JS와 같이 함수는 일급함수로써 다른 함수로부터 리턴되는 리턴값으로 사용 가능하다
func nextValue() func() int {
	i := 0

	//! Closure로 i값 은닉하기
	// i가 nextValue()에서 리턴되지 않는데,
	// i를 리턴하는 익명함수의 상위 렉시컬환경에 있기에 
	// i는 익명함수로만 read되고 modified 될 수 있습니다
	return func() int {
		//* 익명함수가 i를 가지고 있지 않기에
		// 호출된더라도 i의 값은 상위 렉시컬 환경에 state로써 존재한다
		i++
		return i
	}
}

func main() {
	next := nextValue()


	//* 함수 object에서 ()를 붙여야지 함수로써 실행되는 것
	println(next())
	println(next())
	println(next())

	//* 새로운 Closure 환경을 생성하면 i는 0으로 초기화됩니다
	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
}