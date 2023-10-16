package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	success := false
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	answer := rand.Intn(100) + 1 // 1~100
	fmt.Println("Guess number game!")

	for i := 0; i < 10; i++ {
		fmt.Println("Chance left : ", 10-i)
		fmt.Print("Input number : ")
		reader := bufio.NewReader(os.Stdin)
		inputGuessString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputGuessString = strings.TrimSpace(inputGuessString)
		inputGuess, err := strconv.Atoi(inputGuessString)
		if err != nil {
			log.Fatal(err)
		}
		if inputGuess == answer {
			fmt.Println("Great! You got the number, Congraturation~")
			success = true
			break
		} else if inputGuess < answer {
			fmt.Println("Your guess number is lower than answer!")
		} else if inputGuess > answer {
			fmt.Println("Your guess number is higher than answer!")
		}
	}

	if !success {
		fmt.Println("Sorry, you didin't guess my number, It was:", answer)
	}

	/* 	dice := rand.Intn(6) + 1 // 1~6까지 랜덤한 수 출력
	   	fmt.Println(dice) */

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

}
