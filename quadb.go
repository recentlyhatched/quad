package quad

import "fmt"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				// Top-left corner
				fmt.Print("/")
			} else if row == 1 && col == x {
				// Top-right corner
				fmt.Print("\\")
			} else if row == y && col == 1 {
				// Bottom-left corner
				fmt.Print("\\")
			} else if row == y && col == x {
				// Bottom-right corner
				fmt.Print("/")
			} else if row == 1 || row == y {
				// Top and bottom borders
				fmt.Print("*")
			} else if col == 1 || col == x {
				// Left and right borders
				fmt.Print("*")
			} else {
				// Inside
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
