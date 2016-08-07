package main

/* Example #04: https://gobyexample.com/constants */

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
