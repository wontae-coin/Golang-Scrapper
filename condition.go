package main

func main() {
	//* if문의 조건식은 반드시 Boolean이다
	if true {
		println("One")
	}

	//* optional statement
	// if문에서 조건식 사용하기 전에 실행할 수 있음
	// 블록 스코프를 가지고 있다
	i := 6
	max := 10
	if val := i * 2; val < max {
		println(val)
	} else {
		println("2 * val exceeeds 10")
	}

	//* switch
	var name string
	category := 1

	switch category {
	case 1:
		name = "Paper book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

}

// * Go만의 특별한 switch문 용법
// * 1. expression이 없을수 있다 => if elseif처럼 사용할 수 있다. 물록 더욱 간단히!
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No hope")
	}

}
