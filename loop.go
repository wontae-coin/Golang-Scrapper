package main

func main() {
	//* 기본형식: 초기값;조건식;증감
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	//* 조건식만 쓰는 for 루프(다른 언어의 while문)
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)

	//* for문으로 무한루프를 만드려면 초기값;조건식;증감을 모두 생략하면 된다
	// for {
	// 	println("Infinite loop")
	// }

	//* for range(=for in)
	names := []string{"A", "B", "C"}
	for index, name := range names {
		println(index, name)
	}

	//* break, continue, goto
	a := 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // 이후를 실행하지 않고 for 루프 시작으로
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 1 {
		goto END
	}
	println(a)
END:
	println("End")
}
