package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/sethvargo/go-diceware/diceware"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var (
	wrdNumber    int
	passNumber   int
	specialChars = [7]string{"@", "#", "$", "%", "^", "&", "*"}
)

func main() {
	wrdNumber = rand.Intn(10)
	if wrdNumber == 0 {
		wrdNumber += rand.Intn(10)
		return
	}
	userInput()
}

func wordGenerator(wordNumber int, passwdNumber int) {
	for loop := 0; loop < passwdNumber; loop++ {
		list, err := diceware.Generate(wordNumber)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.Join(list, specialChars[rand.Intn(7)]))
	}
}

func userInput() {
	fmt.Println("How many passwords do you need? ")
	_, err := fmt.Scan(&passNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	wordGenerator(wrdNumber, passNumber)
}
