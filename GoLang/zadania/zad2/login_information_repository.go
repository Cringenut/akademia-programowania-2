package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	_ "sync"
)

var accounts sync.Map

func generate_accounts() {
	accounts.Store("login1", "njgrij34")
	accounts.Store("login2", "njgrij34")
}

func login_unique_checker(loginToCheck string) bool {
	unique := true
	accounts.Range(func(key, value interface{}) bool {
		if bytes.Contains([]byte(key.(string)), []byte(loginToCheck)) {
			unique = false
			return false
		}
		return true
	})
	return unique
}

func create_account() {
	var login string
	scanner := bufio.NewScanner(os.Stdin)
login_loop:
	for {
		fmt.Println("Enter login: ")

		scanner.Scan()
		login = scanner.Text() // Assign the value to the outer 'login' variable

		if login_unique_checker(login) {
			break login_loop
		} else {
			fmt.Println("Login already taken")
		}
	}
	fmt.Println("Enter password: ")
	scanner.Scan()
	password := scanner.Text()
	accounts.Store(login, password)
	fmt.Println("Login: " + login + " Password: " + password)
}
