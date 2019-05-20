package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage example:\n\n./combinator a,b x,y,z 1,2,3,4\n\nAuthor: Samuel Lampa")
		return
	}

	args := os.Args[1:]

	sets := make([][]string, len(args))
	for i, arg := range args {
		sets[i] = strings.Split(arg, ",")
	}

	out := combinator(sets...)

	for _, a := range out {
		row := strings.Join(a, ", ")
		fmt.Println(row)
	}
}

func combinator(as ...[]string) [][]string {

	// Default case
	if len(as) == 1 {
		return as
	}

	res := [][]string{}

	if len(as) > 1 {
		bs := combinator(as[1:]...)
		head := as[0]
		for _, p := range head {
			for i := 0; i < len(bs[0]); i++ {
				if len(res) == 0 {
					res = append(res, []string{})
				}
				// Repeat the current part in the current top part of the array
				res[0] = append(res[0], p)
			}
			for j := 0; j < len(bs); j++ {
				if len(res) <= j+1 {
					res = append(res, []string{})
				}
				// Just pad on the rest
				res[j+1] = append(res[j+1], bs[j]...)
			}
		}
	}

	return res
}
