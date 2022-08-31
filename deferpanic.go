package main

import "os"

// * 지연실행 defer
// 함수가 실행 종료되었을 때 추가적인 행동을 명령할 수 있다
func main() {
	f, err := os.Open("Invalid.txt")
	if err != nil {
		//* panic함수
		// 함수를 즉시 멈추고, defer함수들을 모두 실행한 후 즉시 리턴
		// 상위함수에도 똑같이 적용된다(계속 콜스택을 타고 올라가며 적용된다)
		panic(err)
	}
	println("Done")
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))

}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
