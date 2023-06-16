package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func main() {
	fmt.Println("sum:", sum(1, 2))
	fmt.Println("sub:", sub(1, 2))
}
