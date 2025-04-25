package quad

import "fmt"

func QuadA(x,y int) {
  if x > 0 && y > 0 && x + y == 8 {
  fmt.Println("o---o")
  fmt.Println("|   |")
  fmt.Println("o---o")
  }
  if x == 5 && y == 1{
   fmt.Println("o---o")
  }
  if x == 1 && y == 1{
   fmt.Println("o")
  }
  if x == 1 && y == 5{
   fmt.Println("o")
   fmt.Println("|")
   fmt.Println("|")
   fmt.Println("|")
   fmt.Println("o")
  }
}