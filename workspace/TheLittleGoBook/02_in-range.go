package main

import "fmt"
import "math/rand"
import "time"

type Rnge struct {
  min, max int
}

func (r *Rnge) isin(num int) bool {
  return (r.min <= num) && (num <= r.max)
}

func main() {
  rand.Seed(time.Now().UnixNano())
  var num, min int
  var r Rnge
  for i := 1; i <= 10; i++ {
    num = rand.Intn(100)
    min = rand.Intn(100)
    r = Rnge{min: min, max: min + rand.Intn(100)}

    if r.isin(num) {
      fmt.Println("[",r.min,"-",num,"-",r.max,"]")
    } else {
      fmt.Println(num,"is not in [",r.min,"-",r.max,"]")
    }
  }
}
