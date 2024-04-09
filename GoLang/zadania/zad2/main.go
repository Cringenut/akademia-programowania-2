package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	generate_accounts()

loop:
	for {
		fmt.Println("Welcome to account system\n1. Create\n2. Log in\n3. Get accounts\n4. Exit")

		scanner.Scan() // Read user input
		input := scanner.Text()
		option, _ := strconv.Atoi(strings.TrimSpace(input))

		switch option {
		case 1:
			create_account()
		case 2:
		case 3:
		case 4:
			break loop
		default:
			fmt.Println("Choose function from menu")
		}
	}
}
