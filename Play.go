package main

import (
	"fmt"
	"math/rand"
)

var (
	// country that the player has to guess fully to win
	magicCountry string
	guessLetter  string
	guess        = ""
	gameOver     = false
)

func play() {
	magicCountry = Countries[rand.Intn(len(Countries))]
	initGuess()

	for !gameOver {
		_, err := fmt.Scanln(guessLetter)
		if err != nil {
			fmt.Println(InvalidInput)
			return
		}

	}
}

func initGuess() {
	// TODO: Initialize guess here using the length of magicCountry
}
