package main

/* Example #65: https://gobyexample.com/exit */

import "fmt"
import "os"

func main() {
  defer fmt.Println("!")

  os.Exit(3)
}
