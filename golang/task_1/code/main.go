package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kennysong/goeliza"
)

func main() {

	fmt.Println("Eliza: " + goeliza.ElizaHi())

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

		// strip enter char from userInput for windows

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1) // replace it everywhere you find this char
		userInput = strings.Replace(userInput, "\n", "", -1)   // replace return key with nothing for linux based os & macos

		return userInput
	}
}
