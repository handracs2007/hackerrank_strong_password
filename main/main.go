// https://www.hackerrank.com/challenges/strong-password/problem

package main

import (
	"fmt"
	"regexp"
)

func minimumNumber(password string) int32 {
	var toAdd = int32(0)
	var regexDigit, _ = regexp.Compile("[0-9]+")
	var regexLowercase, _ = regexp.Compile("[a-z]+")
	var regexUppercase, _ = regexp.Compile("[A-Z]+")
	var regexSpecial, _ = regexp.Compile("[!@#\\$%\\^&*()\\-\\+]+")

	if !regexDigit.MatchString(password) {
		toAdd++
	}

	if !regexLowercase.MatchString(password) {
		toAdd++
	}

	if !regexUppercase.MatchString(password) {
		toAdd++
	}

	if !regexSpecial.MatchString(password) {
		toAdd++
	}

	if int32(len(password))+toAdd < 6 {
		toAdd += 6 - (int32(len(password)) + toAdd)
	}

	return toAdd
}

func main() {
	for _, chr := range "!@#$%^&*()-+" {
		fmt.Println(minimumNumber("Ab1111" + string(chr)))
	}

	fmt.Println(minimumNumber("#HackerRank"))
}
