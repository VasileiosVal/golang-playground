/*
Simple execution prompting the user to enter a floating number,
then it's converted to an integer and is printed to the screen
*/

package main

import "fmt"

func main() {

	var floatNumber float32
	var intNumber int

	fmt.Printf("Insert a number")
	fmt.Scan(&floatNumber)
	intNumber = int(floatNumber)

	fmt.Printf("The integer is %d", intNumber)
}
