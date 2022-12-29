package main

import (
	"fmt"
	"strconv"
)

func main() {

	inputCheck()

}

func inputCheck() {

	var s string
	var i int

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("->")
	// text, _ := reader.ReadLine()

	fmt.Println("What do you want?")
	for {
		_, err := fmt.Scan(&s)
		i, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println("ERROR: YOU NEED TO ENTER A NUMBER!!!!")
		} else {
			fmt.Println("Got: " + strconv.Itoa(i))
			break
		}
	}
}
