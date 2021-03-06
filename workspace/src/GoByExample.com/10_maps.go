package main

/* Example #10: https://gobyexample.com/maps */

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	v2, prs := m["k1"]
	fmt.Println("prs:", prs, "val:", v2)

	v3, prs := m["k2"]
	fmt.Println("prs:", prs, "val:", v3)

	_, prs3 := m["k2"]
	fmt.Println("prs:", prs3)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
