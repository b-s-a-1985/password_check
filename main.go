package main

import (
	"fmt"
	"log"
	"unicode"
)

func passwordChecker(pw string) bool {

	// Convert the password string into a rune type,
	// which is safe for multi-byte (UTF-8) characters
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	// Define some bool variables for password check
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// Loop over the multi-byte characters one at a time
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	} // for

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// read password from terminal
	var pwd string

	fmt.Print("Enter password: ")
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Fatal(err)
	}

	if passwordChecker(pwd) {
		fmt.Printf("%s is a good password.\n", pwd)
	} else {
		fmt.Printf("%s is a bad password.\n", pwd)
	}
}
