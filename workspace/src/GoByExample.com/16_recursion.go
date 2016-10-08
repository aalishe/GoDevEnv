package main

/* Example #16: https://gobyexample.com/recursion */

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("7! = ", fact(7))
}
