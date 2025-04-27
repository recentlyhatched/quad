package quad

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 || row == y) && (col == 1 || col == x) {
				// Corners
				fmt.Print("o")
			} else if row == 1 || row == y {
				// Top and bottom sides
				fmt.Print("-")
			} else if col == 1 || col == x {
				// Left and right sides
				fmt.Print("|")
			} else {
				// Inside of the rectangle
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
