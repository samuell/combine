package main

import "fmt"

func main() {
	a1 := []string{"a", "b", "c"}
	a2 := []string{"1", "2", "3"}
	a3 := []string{"x", "y", "z"}

	out := combinator(a1, a2, a3)

	fmt.Println("Out: ", out)
}

func combinator(as ...[]string) [][]string {

	// Default case
	if len(as) == 1 {
		return as
	}

	res := [][]string{}

	if len(as) > 1 {
		for _, p := range as[0] {
			for i := 0; i < len(as[1]); i++ {
				if len(res) == 0 {
					res = append(res, []string{})
				}
				// Repeat the current part in the current top part of the array
				res[0] = append(res[0], p)
			}
			for j := 1; j < len(as); j++ {
				if len(res) <= j {
					res = append(res, []string{})
				}
				// Just pad on the rest
				res[j] = append(res[j], as[j]...)
			}
		}
	}

	return res
}
