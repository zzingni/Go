package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)
	inputScore, err := strconv.ParseFloat(inputScoreString, 32)

	var grade string
	if inputScore >= 90 {
		grade = "A grade!"
	} else {
		grade = "under A grade ..."
	}
	fmt.Println("you got", grade)

	/*
		 	Hotspurs := "hm ? j? madi?"
			replacePlayer := strings.NewReplacer("?", "son") // 'Replacer' 타입의 값을 생성하고, 'Replacer' 객체를 반환함.
			player := replacePlayer.Replace(Hotspurs)
			fmt.Println(player)

			var now time.Time
			now = time.Now()
			var year int = now.Year()
			month := now.Month()
			fmt.Println(year, month, now.Day(), now.Hour(), now.Minute(), now.Second())
	*/
}
