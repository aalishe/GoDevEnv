package main

import "fmt"
import "math/rand"
import "time"

func Even(num int) bool {
	return num%2 == 0
}

func Odd(num int) bool {
  return !Even(num)
}

// Generate 10 random numbers and print if it is even or odd

func main() {
	var num int
  rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		num = rand.Intn(100)
		if Even(num) {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
}
