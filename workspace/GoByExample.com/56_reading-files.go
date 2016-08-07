package main

/* Example #56: https://gobyexample.com/reading-files */

import (
  "fmt"
  "bufio"
  "io"
  "io/ioutil"
  "os"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func main() {
  filename := "/tmp/dat"

  dat, err := ioutil.ReadFile(filename)
  check(err)
  fmt.Println(string(dat))

  f, err := os.Open(filename)
  check(err)

  b1 := make([]byte, 5)
  n1, err := f.Read(b1)
  check(err)
  fmt.Printf("1. %d bytes: %s\n", n1, string(b1))

  o2, err := f.Seek(6, 0)
  check(err)
  b2 := make([]byte, 2)
  n2, err := f.Read(b2)
  check(err)
  fmt.Printf("2. %d bytes @ %d: %s\n", n2, o2, string(b2))

  o3, err := f.Seek(6, 0)
  check(err)
  b3 := make([]byte, 2)
  n3, err := io.ReadAtLeast(f, b3, 2)
  check(err)
  fmt.Printf("3. %d bytes @ %d: %s\n", n3, o3, string(b3))

  _, err = f.Seek(0, 0)
  check(err)

  r4 := bufio.NewReader(f)
  b4, err := r4.Peek(5)
  check(err)
  fmt.Printf("4. 5 bytes: %s\n", string(b4))

  f.Close()
}
