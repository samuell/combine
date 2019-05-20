package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print usage
	if len(os.Args) < 2 {
		fmt.Println("Usage example:\n\n./combine a,b x,y,z 1,2,3,4\n\nAuthor: Samuel Lampa")
		return
	}

	// Parse arguments
	programArgs := os.Args[1:]
	input := make([][]string, len(programArgs))
	for i, arg := range programArgs {
		input[i] = strings.Split(arg, ",")
	}

	// Do the trick
	result := combine(input)

	// Print some output
	for _, str := range result {
		row := strings.Join(str, ", ")
		fmt.Println(row)
	}
}

func combine(input [][]string) [][]string {
	// Default case
	if len(input) <= 1 {
		return input
	}

	head := input[0]
	tail := combine(input[1:]) // Recursive call

	result := [][]string{}
	for _, str := range head {
		// Multiply each string in head with the length of the rows in the tail
		// (they are guaranteed to be of equal length)
		for i := 0; i < len(tail[0]); i++ {
			if len(result) == 0 {
				result = append(result, []string{})
			}
			result[0] = append(result[0], str)
		}
		// Multiply the content of each row in the tail with the number of rows
		// in the tail
		for j := 1; j <= len(tail); j++ {
			if len(result) <= j {
				result = append(result, []string{})
			}
			result[j] = append(result[j], tail[j-1]...)
		}
	}

	return result
}
