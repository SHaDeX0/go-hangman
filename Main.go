package main

import "fmt"

import (
	"os"
)

const InvalidInput = "Invalid input provided!"

func main() {
	mainMenu()
}

func mainMenu() {
	for true {
		var option int
		fmt.Println("\n\nHangman" +
			"\n1. Play Game" +
			"\n2. Instructions" +
			"\n3. Exit" +
			"\nEnter your choice: ")

		_, err := fmt.Scanln(&option)
		if nil != err {
			fmt.Println(InvalidInput)
			continue
		}

		switch option {
		case 1:
			play()
		case 2:
			help()
		case 3:
			fmt.Println("Thank you for playing!")
			os.Exit(0)
		default:
			fmt.Println(InvalidInput)
		}
	}
}
