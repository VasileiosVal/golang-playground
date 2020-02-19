/*
Execution prompting the user to enter an integer,
then it adds that number to an int slice, it sorts that slice
and prints it to screen. Execution stops when user types "x".
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := make([]int, 0, 3)
	givenValue := ""

	fmt.Println("Please insert an integer. Press x to exit")
	fmt.Scan(&givenValue)

	for givenValue != "x" {
		num, _ := strconv.Atoi(givenValue)

		arr = append(arr, num)
		sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
		fmt.Println(arr)
		fmt.Println("Please insert an integer. Press x to exit")
		fmt.Scan(&givenValue)
	}

}
