package map

import "fmt"

func LearnMap() {
	//* 해시테이블
	var idMap map[int]string
	//* zero value타입: undefined와 같이 변수가 선언만 되고 초기화가 되지 않았을 때
	//* nil타입: 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value

	//* map을 초기화하는 방법
	idMap = make(map[int]string)
	//* 할당
	idMap[901] = "Apple"
	idMap[134] = "Grape"
	idMap[777] = "Tomato"

	//* 초기화와 할당이 동시에 되는 객체 리터럴
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	//* 값 읽기 (value, isExist)
	val, isExist := tickers["MSFT"]
	if !isExist {
		println("No MSFT ticker")
	}
	println(val)

	//* loop 돌면서 나열하기
	for key, val := range tickers {
		fmt.Println(key, val)
	}
}
