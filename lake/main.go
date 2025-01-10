//filename: main.go
// main.go는 컴파일 하기 위한 func (only for compile)

package main

// package import
import (
	"fmt"

	"github.com/something"
)

// filename으로 func 정의
// Go는 Python이나 Node.js와는 다르게 특정 function을 찾게 되는데 그것이 func main()이다.
// 이것이 Go의 시작.
// 자동적으로 컴파일러는 main package와 main function을 찾고, 실행
// main.go는 무조건 compile을 위한 것.
func main_1() {

	// fmt: fmt func를 import하고 Println을 가져와 사용
	fmt.Println("Hello, Lake!")

	//public 함수는 다른 패키지에서 사용 가능
	something.SayHello()

	// private 함수는 다른 패키지에서 사용 불가능
	// something.sayBye()

	// Go는 정적 타입 언어이기 때문에 변수를 선언할 때 타입을 지정해주어야 한다.
	const name string = "Lake"

	// name을 bloon으로 설정했는데 값이 String이라 오류 발생
	// const name bloon = "Lake"

	// name은 const로 선언되어 있기 때문에 변경 불가능
	// name을 변경하려고 하면 오류 발생
	// name = "henrry"
	fmt.Println(name)

	var nickanme string = "Lake"

	// var로 선언된 변수는 변경 가능
	nickanme = "henrry"
	fmt.Println(nickanme)

	// type이 없는 변수는 :=로 선언
	// 축약형은 오직 함수(func) 내에서만 사용 가능!!
	host := "Lake"
	fmt.Println(host)

	// go가 할당받은 값을 통해 type을 추론하여 host2를 bloon으로 설정
	host2 := false
	fmt.Println(host2)

}
