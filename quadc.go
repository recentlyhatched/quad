package quad

import (
	"fmt"
)

func QuadC(x,y int) {
	// first loop - rows A-B-B-etc-C
	for i := 0; i < y; i++ {
		if i == 0 {
			// first letter - left-most 'A'
			fmt.Print("A")
		} else if i == y-1 {
			// left-most bottom 'C'
			fmt.Print("C")
		} else {
			// left-most middle 'B's
			fmt.Print("B")
		}
		// second loop - coloumns - After  starting ABBBetcA/C
		for j := 0; j < x-1; j++ {
			if i == 0 && j == x-2 {
				// ABBB'A'
				fmt.Print("A")
			} else if i == y-1 && j == x-2 {
				// last letter - CBBB'C'
				fmt.Print("C")
			} else if j == x-2 || i == 0 || i == y-1 {
				// right-most outer column 'B'
				fmt.Print("B")
			} else {
				// middle characters
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}