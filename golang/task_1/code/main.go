package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kennysong/goeliza"
)

func main() {

	fmt.Println("Eliza: " + goeliza.ElizaHi())

	fmt.Println("UserInput: " + getUserInput())

	for {
		statement := getUserInput()
		fmt.Println("Eliza: " + goeliza.ReplyTo(statement))

		if goeliza.IsQuitStatement(statement) {
			break
		}
	}

}

func getUserInput() string {

	for {
		// get reader input
		reader := bufio.NewReader(os.Stdin)

		userInput, _ := reader.ReadString('\n')

		fmt.Println(userInput)
	}
}
