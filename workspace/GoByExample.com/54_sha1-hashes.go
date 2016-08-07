package main

/* Example #54: https://gobyexample.com/sha1-hashes */

import "fmt"
import "crypto/sha1"
import "crypto/md5"

func main() {
  s := "sha1 this string"
  h := sha1.New()
  h.Write([]byte(s))
  bs := h.Sum(nil)

  fmt.Println(s)
  fmt.Printf("SHA1: %x\n", bs)

  m := md5.New()
  m.Write([]byte(s))
  bms := m.Sum(nil)

  fmt.Println(s)
  fmt.Printf("MD5: %x\n", bms)
}
