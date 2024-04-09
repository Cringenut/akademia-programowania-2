package main

import (
	"bytes"
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
