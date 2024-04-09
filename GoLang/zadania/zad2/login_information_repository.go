package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	_ "sync"
)

var accounts sync.Map

func generate_accounts() {
	accounts.Store("login1", "njgrij34")
	accounts.Store("login2", "njgrij34")
}

func contains_login(loginToCheck string) bool {
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
	var builder strings.Builder
login_loop:
	for {
		fmt.Println("Enter login: ")

		scanner.Scan()
		login = scanner.Text() // Assign the value to the outer 'login' variable

		if contains_login(login) {
			break login_loop
		} else {
			fmt.Println("Login already taken")
		}
	}
	fmt.Println("Enter password: ")
	scanner.Scan()
	password := scanner.Text()
	accounts.Store(login, password)

	fmt.Fprintf(&builder, "Login: %s Password: %s\n", login, password)
	fmt.Println(builder.String())
}

func print_accounts() {
	var builder strings.Builder

	accounts.Range(func(key, value interface{}) bool {
		login := key.(string)
		password := value.(string)
		fmt.Fprintf(&builder, "Login: %s, Password: %s\n", login, password)
		return true // continue iterating
	})

	fmt.Println(builder.String())
}
