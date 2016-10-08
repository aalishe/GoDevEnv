package main

import "fmt"

func main()  {
  var x [5]int
  x[4] = 100
  fmt.Println(x)

  y := [5]float64{
    98,
    93,
    84,
    64,
    95,
  }
  var total float64 = 0
  for i := 0; i < len(y); i++ {
    total += y[i]
  }
  fmt.Println("Average: ", total / float64(len(y)))

  total = 0
  for _, val := range y {
    total += val
  }
  fmt.Println("Average: ", total / float64(len(y)))

  z := make([]float64, 5)
  v := [5]float64{1,2,3,4,5}
  fmt.Println("Z: ", z)
  fmt.Println("V: ", v)

  a := make(map[string]int)
  a["key"] = 10
  fmt.Println("A: ", a)

  elements := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be": "Beryllium",
    "B":  "Boron",
    "C":  "Carbon",
    "N":  "Nitrogen",
    "O":  "Oxygen",
    "F":  "Fluorine",
    "Ne": "Neon",
  }

  name, ok := elements["Un"]
  fmt.Println("Name: ",name, "OK: ",ok)
  if n,ok := elements["H"]; ok {
    fmt.Println("Name: ", n, "OK: ",ok)
  }

  complex := map[string]map[string]string{
    "H": map[string]string{
      "name": "Hydrogen",
      "state": "gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map[string]string{
      "name":"Beryllium",
      "state":"solid",
    },
    "B": map[string]string{
      "name":"Boron",
      "state":"solid",
    },
    "C": map[string]string{
      "name":"Carbon",
      "state":"solid",
    },
    "N": map[string]string{
      "name":"Nitrogen",
      "state":"gas",
    },
    "O": map[string]string{
      "name":"Oxygen",
      "state":"gas",
    },
    "F": map[string]string{
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne": map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }

  if el, ok := complex["Li"]; ok {
    fmt.Println("Name: ", el["name"], "State: ", el["state"])
  }

  lst := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  min := 10000
  index := -1
  for i, val := range lst {
    if val < min {
      min = val
      index = i
    }
  }
  fmt.Println("The minimun is: ", lst[index])

}
