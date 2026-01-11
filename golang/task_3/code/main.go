package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)

	}
	defer func() {
		_ = keyboard.Close()
	}()

	// Need to restructure this code

	// 1. Welcome Users with a Nice Prompt to your Restaurant
	// 2. As them to provide their names, email and locations (store in a local cache database)
	// 3. Ask them what they'd like help with today (Give them options) - Maybe 3
	// 4. Expand the catalogue to show whats on the Menu after selection
	// 5. Once they order, process the order and send them an email
	// 6. Thank them for their order
	// 7. Make sure you handle cases where users choose wrong[ non existent ] menu on item etc

	fmt.Println("Welcome to Nana's Restaurant")
	fmt.Println("Please Choose Your Order from the menu")
	fmt.Println("Press ESC to exit")

	characterPressed, keyPressed := GetUserPrompt()

	fmt.Printf("You've pressed : %s %s", characterPressed, keyPressed)

	// using rune here instead of int as rune is already a char
	// and doesn't need to be converted to string NOT rune is
	// written '1' with just ''
	menu := map[rune]string{
		'1': "Cappucino",
		'2': "Americano",
		'3': "Latte",
		'4': "Latte Macciato",
		'5': "Mocha",
		'6': "Caramel Macciato",
	}

	// store just the keys to be sliced and sorted later
	key := []rune{'1', '2', '3', '4', '5', '6'}

	// using the sort function take a slice of this key
	// and run this function on it
	// its sorts it. Logic if the value on the left of
	// the key is less than the one on the right return true.
	// does this for all keys until its sorted in ascending order
	// is key[j] > k[i] for decending order
	sort.Slice(key, func(i, j int) bool {
		return key[i] < key[j]

	})

	for _, k := range key {
		// char sub chars %s returns strings, %d for ints %v prints out the output in its original/whatever
		//format
		fmt.Printf("%s %s\n", string(k), menu[k])
	}

	fmt.Println("\n Press A Number to Order or ESC to exit")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error : %v\n", err)
			break
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)

		if key == keyboard.KeyEsc {
			fmt.Println("---- Thanks for passing Through : Goodbye -----")
			break
		}

		// get the menu using the character entered
		// it shows the menu[2] so item on menu & true
		// it returns the item in the list and true if the character
		// entered exists in the map
		// it returns %s % false if the char doesn't exists in the map
		order, exists := menu[char]

		if !exists {
			fmt.Println("What you selected isn't in the menu")
			break
		} else {
			return
		}

		fmt.Printf("%s %d", order, exists)

		fmt.Println("----------")

	}

}

func GetUserPrompt() (rune, keyboard.Key) {

	for {
		char, key, error := keyboard.GetKey()
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error : %v\n", error)
			return 0, 0
		}

		fmt.Printf("You've pressed : %s %s", char, key)
		return char, key
	}

}

func CheckInventory(r rune) (v int, s bool) {

	if r > '0' && r <= '6' {
		val, _ := strconv.Atoi(string(r))

		return val, true
	} else {

		GetUserPrompt()

	}

}
