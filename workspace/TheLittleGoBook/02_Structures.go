package main

import "fmt"

type Saiyan struct {
  Name string
  Power int
}

func (s *Saiyan) Super()  {
  s.Power += 10000
}

func NewSaiyan(name string, power int) *Saiyan {
  return &Saiyan{
    Name: name,
    Power: power,
  }
}

type Person struct {
  Name string
}

func (p *Person) Introduce()  {
  fmt.Printf("Hi, I am %s\n", p.Name)
}

type Saiyan2 struct {
  *Person
  Power int
}

func (s *Saiyan2) Introduce() {
  fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

func main() {
  goku := Saiyan {
    Name: "Goku",
    Power: 900,
  }
  fmt.Println(goku)

  goku2 := new(Saiyan)
  // Same as:
  // goku2 := Saiyan{}
  goku2.Name = "Goku"
  goku2.Power = 9000
  fmt.Println(goku2)

  goku3 := &Saiyan{"Goku", 9000}
  goku3.Super()
  fmt.Println(goku3.Power)

  goku4 := &Saiyan2{
    Person: &Person{"Goku"},
    Power: 9001,
  }
  goku4.Introduce()
  goku4.Person.Introduce()
  fmt.Println(goku4.Name, "==", goku4.Person.Name)
}
