package main

/* Example #60: https://gobyexample.com/command-line-flags */

import "fmt"
import "flag"

func main() {
  wordPtr := flag.String("word", "foo", "a string")
  numbPtr := flag.Int("numb", 42, "an int")
  boolPtr := flag.Bool("fork", false, "a bool")

  var svar string
  flag.StringVar(&svar, "svar", "bar", "a string var")

  flag.Parse()

  fmt.Println("Word:", *wordPtr)
  fmt.Println("Numb:", *numbPtr)
  fmt.Println("Fork:", *boolPtr)
  fmt.Println("sVar:", svar)
  fmt.Println("Tail:", flag.Args())
}
