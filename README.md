# Quad

## Description
This project is written in Go. It runs based on the calling QuadA(), QuadB(), QuadC(), QuadD() or QuadE(). In Golang, Pascal case is used to export functions. These functions take two numbers as parameters. The first parameter represents the width of the output and the second parameter represents the length.

Example: QuadA(5, 1)\
Terminal output: o---o

## Authors
Georgina, Masoud, Anatoly, Redder

## Usage
1. Make sure you have the latest version of Go\
Go to your terminal and run the following:
2. `git clone <repo-link>`
3. `cd quad`
4. In the test/main.go file, change parameters in QuadA(x, y) where x is the width and y is the height, or use other functions defined
4. `go run test/main.go`

## Implementation details
To write a function that prints a valid retangle, we used for-loops and conditional if-statements. With the for loop, we were able to iterate through each column and row using nested loops for basic/repeat characters. Then, based on which column or row, print special "corner" characters.

### Examples:
Function: quad.QuadC(3, 6)\
Output:\
`ABA
B B
B B
B B
B B
CBC`

Function: quad.QuadE(7, 4)\
Output:\
`ABBBBBC
B     B
B     B
CBBBBBA`