package main

import (
	"fmt"
	"strconv"
)

func help() {
	fmt.Println("I will pick a random country name.\n" +
		"That is the Magic Country which you will have to guess.\n" +
		"You will have " + strconv.Itoa(Lives) + " lives to completely guessCountry the Magic Country.\n" +
		"If you guessCountry a letter which is not present in the Magic Country, you will lose a life.\n" +
		"When you lose all your lives or guessCountry the Magic Country name, the game will be over.\n" +
		"Now, let's get started!\n\n")
}
