package quad

import (
	"fmt"
)

func QuadC(x,y int) {
	if x > 0 && y > 0 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				fmt.Print("A")
			} else {
				fmt.Print("B")
			}
			if(i > 0 && i < x-1) {
				fmt.Print("\n")
				for j := 0; j < y; j++ {
					fmt.Print("B")
				}
			}
		}
	}
}