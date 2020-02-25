/*
Execution prompting the user to enter a name of a local file, containing
the first and last name of persons. The process iterrates through each line,
reads the data, creates a struct which takes those two params and stores that
struct into a slice. Finally it iterrates through that slice and prints every struct.

e.g. of that sample file (names.txt)
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	fname, lname string
}

func main() {
	var fileName string
	var personSlice []Person

	fmt.Println("Please provide the name of the file")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error", err)
	} else {
		defer file.Close()
		data := bufio.NewReader(file)
		for {
			line, err := data.ReadString('\n')

			person := strings.Split(line, " ")
			personStruct := Person{person[0], person[1]}
			personSlice = append(personSlice, personStruct)
			if err == io.EOF {
				fmt.Println("finished reading file")
				break
			}
			if err != nil {
				fmt.Println("error", err)
			}
		}
		fmt.Println("Printing every person's data:")
		for i, v := range personSlice {
			fmt.Printf("%d) firstName: %v - lastName: %v\n", i+1, v.fname, v.lname)
		}
	}
}
