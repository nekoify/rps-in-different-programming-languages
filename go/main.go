package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func main()  {
	fmt.Println("Welcome to Rock Paper Sissors written in Go")
	rps()
}

func rps() {
	var userInput string
	fmt.Println("Please select Rock, Paper or Scissors: ")
	fmt.Scanln(&userInput)
	var answer string = strings.ToLower(userInput)
	var res bool = checkAnswer(answer)
	if res == false {return}
	var aiAnswer int = rand.Intn(3 - 1) + 1
	var aiMove = parseAnswer(aiAnswer)
	fmt.Println(fmt.Sprintf("Ai has picked %s", aiMove))
	fmt.Println(findWinner(answer, aiMove))
}

func checkAnswer(ans string)(bool)  {
	if !(ans == "rock" || ans == "paper" || ans == "scissors") {
		fmt.Println("Please enter a valid move!")
		rps()
		return false 
	}
	return true
}

func parseAnswer(ans int)(string)  {
	if ans == 1 {
		return "rock"
	} else if ans == 2 {
		return "paper"
	} else if (ans == 3) {
		return "scissors"
	}
	return "?"
}

func findWinner(user string, ai string)(string)  {
	if (user == "rock" && ai == "scissors") ||
		(user == "paper" && ai == "rock") ||
		(user == "scissors" && ai == "paper") {
		return "You win!"
	} else if (ai == "rock" && user == "scissors") ||
		(ai == "paper" && user == "rock") ||
		(ai == "scissors" && user == "paper") {
		return "The AI wins!"
	} else {
		return "It's a tie!"
	}
}