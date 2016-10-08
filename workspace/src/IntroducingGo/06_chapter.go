package main

import "fmt"

func average(xs []float64) float64 {
  var total float64 = 0.0
  for i := 0; i < len(xs); i++ {
    total += xs[i]
  }
  return total / float64(len(xs))
}

func main()  {
  xs := []float64{98, 93, 90, 23, 53}

  ave := average(xs)
  fmt.Println("Average: ", ave)
}
