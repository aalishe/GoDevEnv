package main

import "fmt"

func main()  {
  for i := 1; i <= 10; i++ {
    fmt.Print(i,": ")
    switch i {
    case 0: fmt.Print(" (Zero) ")
    case 1: fmt.Print(" (One) ")
    case 2: fmt.Print(" (Two) ")
    case 3: fmt.Print(" (Three) ")
    case 4: fmt.Print(" (Four) ")
    case 5: fmt.Print(" (Five) ")
    case 6: fmt.Print(" (Six) ")
    case 7: fmt.Print(" (Seven) ")
    case 8: fmt.Print(" (Eight) ")
    case 9: fmt.Print(" (Nine) ")
    case 10: fmt.Print(" (Ten) ")
    default: fmt.Println(" (Unknown) ")
    }
    if i % 2 == 0 {
      fmt.Println("even")
    } else {
      fmt.Println("odd")
    }
  }

  for j := 0; j <= 100; j++ {
    if j % 3 == 0 {
      fmt.Print(j,", ")
    }
  }
}
