package main

/* Example #58: https://gobyexample.com/line-filters */

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    ucl := strings.ToUpper(scanner.Text())
    fmt.Println(ucl)
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
  }
}
