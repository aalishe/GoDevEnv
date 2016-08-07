package main

/* Example #41: https://gobyexample.com/panic */

import "os"

func main() {
  panic("a problem")

  _, err := os.Create("/tmp/file")
  if err != nil {
    panic(err)
  }
}
