package quad

import (
	"fmt"
)

func QuadC(x,y int) {
	if (x > 0 && y > 0) {
		for i : 0; i < x; i++ {
			fmt.Print("B")
		}
	}
}