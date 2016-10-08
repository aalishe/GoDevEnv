package main

import "fmt"

func main()  {
  var x string = "Hi"
  fmt.Println(x)
  var (
    a = 5
    b = 10
    d = 15
  )
  fmt.Println("Sum: ", a + b + d)
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  fmt.Println("Double: ", input * 2)

  fmt.Print("Enter Fahrenheit: ")
  var f float64
  fmt.Scanf("%f", &f)
  c := (f - 32) * 5/9
  fmt.Println("Celsius: ", c)

  fmt.Print("Enter Feets: ")
  var fts float64
  fmt.Scanf("%f", &fts)
  m := f * 0.3048
  fmt.Println("Meters: ", m)
}
