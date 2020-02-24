/*
Execution prompting the user to enter a name and an address (strings),
then it creates a map with those two as keys (name, address), filled
with the values the user entered. Then it converts that map to JSON
and prints it's encoded ([]byte) and decoded output (JSON object)
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name, address string

	fmt.Println("Please enter your name:")
	fmt.Scan(&name)
	fmt.Println("Please enter your address:")
	fmt.Scan(&address)

	credits := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonCredits, err := json.Marshal(credits)
	if err != nil {
		fmt.Println("error:", err)
	} else {

		fmt.Println("The encoded JSON output is:", jsonCredits)
		fmt.Println("The decoded JSON output is:", string(jsonCredits))
	}

}
