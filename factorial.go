package main

import (
	"fmt"
	"log"
)

func fact(n int) int {

	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}

func main() {
	var in int

	fmt.Print("정수 입력: ")
	_, err := fmt.Scanln(&in)

	if err != nil {
		log.Fatal(err)
	}

	result := fact(in)
	fmt.Printf("%d!는 %d입니다.", in, result)

}
