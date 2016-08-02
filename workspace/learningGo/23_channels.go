package main

/* Example #23: https://gobyexample.com/channels */

import "fmt"

func main() {
  messages := make(chan string)

  go func()  { messages <- "ping" }()

  msg := <-messages
  fmt.Println(msg)
}
