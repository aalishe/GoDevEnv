package main

/* Example #64: https://gobyexample.com/signals */

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)

  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

  go func() {
    sig := <-sigs
    fmt.Println()
    fmt.Println(sig)
    done <- true
  }()

  fmt.Println("Awaiting signal")
  fmt.Println("From another terminal kill the process with -2 and -15 (default)")
  <-done
  fmt.Println("Exiting")
}
