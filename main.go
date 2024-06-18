package main

import (
	"fmt"
	"unicode"
)




func passwordChecker(pw string) bool {

	// Convert the password string into a rune type,
	// which is safe for multi-byte (UTF-8) characters
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	// Define some bool variables 
	hasUpper := false	
	hasLower := false
	hasNumber := false
	hasSymbol := false

	

	// Loop over the multi-byte characters one at a time
	for _, v := range pwR{
		if unicode.IsUpper(v){
			hasUpper = true
		}
		if unicode.IsLower(v){
			hasLower = true
		}
		if unicode.IsNumber(v){
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v){
			hasSymbol = true
		}
	} // for

	return hasUpper && hasLower && hasNumber && hasSymbol

}

func main(){
	pwd := "Aa_12345"
	if passwordChecker(pwd){
		fmt.Println("good password")
	} else {
		fmt.Println("bad password")
	}
	
}