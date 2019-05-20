package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage example:\n\n./combine a,b x,y,z 1,2,3,4\n\nAuthor: Samuel Lampa")
		return
	}

	programArgs := os.Args[1:]

	input := make([][]string, len(programArgs))
	for i, arg := range programArgs {
		input[i] = strings.Split(arg, ",")
	}

	result := combine(input)

	for _, str := range result {
		row := strings.Join(str, ", ")
		fmt.Println(row)
	}
}

func combine(input [][]string) [][]string {

	// Default case
	if len(input) == 1 {
		return input
	}

	result := [][]string{}

	if len(input) > 1 {
		head := input[0]
		tail := combine(input[1:]) // Recursive call
		for _, p := range head {
			for i := 0; i < len(tail[0]); i++ {
				if len(result) == 0 {
					result = append(result, []string{})
				}
				// Repeat the current part in the current top part of the array
				result[0] = append(result[0], p)
			}
			for j := 1; j <= len(tail); j++ {
				if len(result) <= j {
					result = append(result, []string{})
				}
				// Just pad on the rest
				result[j] = append(result[j], tail[j-1]...)
			}
		}
	}

	return result
}
