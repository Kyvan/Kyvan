package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sethvargo/go-diceware/diceware"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var (
	userArg      string
	wrdNumber    int
	passNumber   int
	min          int = 2
	max          int = 10
	specialChars     = [7]string{"@", "#", "$", "%", "^", "&", "*"}
	password     string
)

func inputChecker() {
	wrdNumber = rand.Intn((max - min + 1) + min)
	if wrdNumber <= 2 {
		wrdNumber += rand.Intn(10)
		return
	}

	if len(os.Args) == 1 {
		userInput()
	} else if len(os.Args) == 2 {
		userArg = os.Args[1]
		userNumbers, err := strconv.Atoi(userArg)
		if err == nil {
			wordGenerator(userNumbers, wrdNumber)
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
		password = strings.Join(list, specialChars[rand.Intn(7)])
		println(password)
	}
}

func userInput() {
	fmt.Println("How many passwords do you need? ")
	_, err := fmt.Scan(&passNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	wordGenerator(passNumber, wrdNumber)
}

func main() {
	inputChecker()
}
