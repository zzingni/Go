package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	var b float32 = 8.34
	fmt.Println(a * int(b))
	var c7 string // 변수명은 알파벳 대소문자로 시작해야 한다.
	var d int
	var e bool
	var f float64
	var G = 99

	koreanZombie := "정찬성"
	fmt.Println(koreanZombie)

	fmt.Println(c7, d, e, f, G) // 문자열에는 값을 담지 않으면 아무것도 출력 안됨. int, float64는 0, bool값은 flase!
	fmt.Println(b, reflect.TypeOf(b))

	// fmt.Println(math.Floor("삼.오"), math.Ceil("이백십칠쩜칠"), math.Sqrt("sixteen"))
	fmt.Println(math.Floor(2.75))
	fmt.Println(math.Ceil(2.75))
	fmt.Println(math.Sqrt(16))
	fmt.Println("Opensource\tprogramming\n\"go!\"")
	fmt.Println(strings.Title("head first go"))
}
