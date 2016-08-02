package main

/* Example #31: https://gobyexample.com/range-over-channels */

import "fmt"

func main() {
  queue := make(chan string, 2)
  queue <- "one"
  queue <- "two"
  close(queue)

  for elem := range queue {
    fmt.Println(elem)
  }
}
