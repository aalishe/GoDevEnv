package main

/* Example #33: https://gobyexample.com/tickers */

import "fmt"
import "time"

func main() {
  ticker := time.NewTicker(time.Millisecond * 500)
  go func ()  {
    for t := range ticker.C {
      fmt.Println("Tick at", t)
    }
  }()

  time.Sleep(time.Millisecond * 2600)
  ticker.Stop()
  fmt.Println("Ticker stopped")
}
