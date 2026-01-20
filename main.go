package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		user_input := cleanInput(scanner.Text())
		fmt.Println("Your command was: " + user_input[0])
	}
}

