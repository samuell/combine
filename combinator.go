package main

import "fmt"

func main() {
	a1 := []string{"a", "b", "c"}
	a2 := []string{"1", "2", "3"}
	a3 := []string{"x", "y", "z"}

	out := combinator(a1, a2, a3)

	for i, a := range out {
		fmt.Printf("Out array [%d]: %v\n", i, a)
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
		for _, p := range as[0] {
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
				res[j+1] = append(res[j], bs[j]...)
			}
		}
	}

	return res
}
