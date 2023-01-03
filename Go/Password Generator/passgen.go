package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
)

var (
	_wrdNumber   int
	userArg      string
	passNumber   int
	specialChars = []string{"@", "#", "$", "%", "^", "&", "*", "!", "-", "="}
	specialNumbs = []string{"2", "3", "4", "5", "6", "7", "8", "9", "69", "420"}
	password     string
)

type IntRange struct {
	min, max int
}

func (wrdNumber *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(wrdNumber.max-wrdNumber.min+1) + wrdNumber.min
}

func randomNumber(_userNumbers int) {
	r := rand.New(rand.NewSource(10))
	wrdNumber := IntRange{1, 5}
	_wrdNumber = wrdNumber.NextRandom(r)
	wordGenerator(_userNumbers, _wrdNumber)
}

func inputChecker() {
	if len(os.Args) == 1 {
		userInput()
	} else if len(os.Args) == 2 {

		userArg = os.Args[1]
		userNumbers, err := strconv.Atoi(userArg)
		if err == nil {
			randomNumber(userNumbers)
		}
	} else {
		log.Fatal("ERROR: You need either 0 or 1 arguments")
	}
}

func wordGenerator(passwdNumber int, wordNumber int) {
	for loop := 0; loop < passwdNumber; loop++ {
		list, err := diceware.Generate(wordNumber)
		if err != nil {
			log.Fatal(err)
		}
		password = strings.Join(list, specialChars[rand.Intn(len(specialChars))]+specialNumbs[rand.Intn(len(specialNumbs))])
		fmt.Println(password)

	}
}

func userInput() {
	fmt.Println("How many passwords do you need?")
	_, err := fmt.Scan(&passNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	randomNumber(passNumber)
}

func main() {
	inputChecker()
}
