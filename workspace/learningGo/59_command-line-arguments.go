package main

/* Example #59: https://gobyexample.com/command-line-arguments */

import "fmt"
import "os"

func main() {
  argsWithProg := os.Args
  argsWithoutProg := os.Args[1:]

  arg := os.Args[3]

  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)
  fmt.Println(arg)
}
