package main

/* Example #57: https://gobyexample.com/writing-files */

import (
  "fmt"
  "bufio"
  "io/ioutil"
  "os"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func main() {
  d1 := []byte("hello\ngo\n")
  err := ioutil.WriteFile("/tmp/dat", d1, 0644)
  check(err)

  f, err := os.Create("/tmp/dat1")
  check(err)

  defer f.Close()

  d2 := []byte{115, 111, 109, 101, 10}
  n2, err := f.Write(d2)
  check(err)
  fmt.Printf("1. Wrote %d bytes\n", n2)

  n3, err := f.WriteString("writes\n")
  fmt.Printf("2. Wrote %d bytes\n", n3)

  f.Sync()

  w := bufio.NewWriter(f)
  n4, err := w.WriteString("buffered\n")
  fmt.Printf("3. Wrote %d bytes\n", n4)

  w.Flush()
}
