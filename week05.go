package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	Hotspurs := "hm ? j? madi?"
	replacePlayer := strings.NewReplacer("?", "son") // 'Replacer' 타입의 값을 생성하고, 'Replacer' 객체를 반환함.
	player := replacePlayer.Replace(Hotspurs)
	fmt.Println(player)

	var now time.Time
	now = time.Now()
	var year int = now.Year()
	month := now.Month()
	fmt.Println(year, month, now.Day(), now.Hour(), now.Minute(), now.Second())
}
