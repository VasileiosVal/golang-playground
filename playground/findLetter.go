/*
Execution prompting the user to enter a string,
then it checks if the given string starts with letter "a",
ends with letter "c" and it contains letter "s".
Finally it prints the result of that check.
*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var givenString string

	fmt.Println("Insert a string: ")
	fmt.Scan(&givenString)

	var firstLetterFound = strings.HasPrefix(strings.ToLower(givenString), "a")
	var lastLetterFound = strings.HasSuffix(strings.ToLower(givenString), "s")
	var middleLetterPosition = strings.Index(strings.ToLower(givenString), "c")
	var givenStringLength = utf8.RuneCountInString(givenString)

	if firstLetterFound && lastLetterFound && middleLetterPosition > 0 && middleLetterPosition < givenStringLength-1 {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
