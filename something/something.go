package something

import "fmt"

// 앞글자가 대문자: public 함수
// public 함수는 다른 패키지에서 사용 가능
func SayHello() string {
	fmt.Println("Hello!")
}

// 앞글자가 소문자: private 함수
// private 함수는 다른 패키지에서 사용 불가능
func sayBye() string {
	fmt.Println("Bye!")
}
