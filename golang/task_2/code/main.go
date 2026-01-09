package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "And wait please"

func main() {
	// seed the random number generator to always give a random number
	rand.Seed(time.Now().UnixNano())

	// generate a random numbe between 0 - 8 and add two
	first_num := rand.Intn(8) + 2
	second_num := rand.Intn(8) + 2
	subtraction := rand.Intn(8) + 2
	answer := first_num*second_num - subtraction

	userInput(first_num, second_num, subtraction, answer)

}

func userInput(first_num, second_num, subtraction, answer int) {

	//init the userdata input stream
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("-----------------------")
	fmt.Println(" ")

	fmt.Println("Think of number between 0 - 10")
	reader.ReadString('\n')

	fmt.Println("Multiple your number by:", first_num, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiple the results by", second_num, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the results by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	fmt.Println("The number you took was", answer)

}
