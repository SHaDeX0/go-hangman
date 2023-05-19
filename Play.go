package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	countries    = [...]string{"Afghanistan", "Albania", "Algeria", "Andorra", "Angola", "Anguilla", "Antigua Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia Herzegovina", "Botswana", "Brazil", "British Virgin Islands", "Brunei", "Bulgaria", "Burkina Faso", "Burundi", "Cambodia", "Cameroon", "Cape Verde", "Cayman Islands", "Chad", "Chile", "China", "Colombia", "Congo", "Cook Islands", "Costa Rica", "Cote D Ivoire", "Croatia", "Cruise Ship", "Cuba", "Cyprus", "Czech Republic", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador", "Equatorial Guinea", "Estonia", "Ethiopia", "Falkland Islands", "Faroe Islands", "Fiji", "Finland", "France", "French Polynesia", "French West Indies", "Gabon", "Gambia", "Georgia", "Germany", "Ghana", "Gibraltar", "Greece", "Greenland", "Grenada", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea Bissau", "Guyana", "Haiti", "Honduras", "Hong Kong", "Hungary", "Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy", "Jamaica", "Japan", "Jersey", "Jordan", "Kazakhstan", "Kenya", "Kuwait", "Kyrgyz Republic", "Laos", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania", "Luxembourg", "Macau", "Macedonia", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Mauritania", "Mauritius", "Mexico", "Moldova", "Monaco", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Namibia", "Nepal", "Netherlands", "Netherlands Antilles", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria", "Norway", "Oman", "Pakistan", "Palestine", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Poland", "Portugal", "Puerto Rico", "Qatar", "Reunion", "Romania", "Russia", "Rwanda", "Saint Pierre Miquelon", "Samoa", "San Marino", "Satellite", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovakia", "Slovenia", "South Africa", "South Korea", "Spain", "Sri Lanka", "St Kitts Nevis", "St Lucia", "St Vincent", "St. Lucia", "Sudan", "Suriname", "Swaziland", "Sweden", "Switzerland", "Syria", "Taiwan", "Tajikistan", "Tanzania", "Thailand", "Timor L'Este", "Togo", "Tonga", "Trinidad Tobago", "Tunisia", "Turkey", "Turkmenistan", "Turks Caicos", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom", "Uruguay", "Uzbekistan", "Venezuela", "Vietnam", "Virgin Islands (US)", "Yemen", "Zambia", "Zimbabwe"}
	gameOver     bool
	Lives        int
	magicCountry string
	guessLetter  string
	guessCountry string
	wrongGuesses string
)

func play() {
	initPlay()

	for !gameOver {
		fmt.Println("\n\nLives remaining: ", Lives)
		fmt.Println("Country guess: " + guessCountry)
		fmt.Println("\n\nEnter your guess letter: ")
		_, err := fmt.Scanln(&guessLetter)
		if err != nil {
			fmt.Println(InvalidInput)
			continue
		}

		if 1 < len(guessLetter) {
			fmt.Println("Enter just a single letter!")
			continue
		}

		if !(strings.Contains(magicCountry, guessLetter)) {
			checkForWrongGuesses()
			checkForLose()
			continue
		}

		modifyGuess()
		checkForWin()
	}

	if 0 == Lives {
		fmt.Println("You lost! The Magic Country name was: " + magicCountry)
	} else {
		fmt.Println("Congratulations! You guessed the Magic Country: " + magicCountry)
	}
}

func checkForWrongGuesses() {
	if strings.Contains(wrongGuesses, guessLetter) {
		fmt.Println("You already guessed this letter, and it's wrong!")
		return
	}
	fmt.Println("Wrong guess!")
	wrongGuesses += guessLetter
	Lives--
}

func checkForLose() {
	if 0 == Lives {
		gameOver = true
	}
}

func checkForWin() {
	win := true
	for i := 0; i < len(magicCountry); i++ {
		if guessCountry[i] != magicCountry[i] {
			win = false
			break
		}
	}
	if win {
		gameOver = true
	}
}

func initPlay() {
	fmt.Println("Countries length: ", len(countries))
	magicCountry = strings.ToLower(countries[rand.Intn(len(countries))])
	gameOver = false
	guessCountry = ""
	wrongGuesses = ""
	Lives = 10
	// Initialize guessCountry here using the length of magicCountry
	for i := 0; i < len(magicCountry); i++ {
		if 32 == magicCountry[i] {
			guessCountry += " "
			continue
		}
		guessCountry += "_"
	}
}

func modifyGuess() {
	for i := 0; i < len(magicCountry); i++ {
		if magicCountry[i] == guessLetter[0] {
			guessCountry = guessCountry[:i] + guessLetter + guessCountry[i+1:]
		}
	}
}
