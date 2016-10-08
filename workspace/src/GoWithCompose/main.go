package main

import (
  "fmt"
  redis "gopkg.in/redis.v4"
)

func main()  {
  client := redis.NewClient(&redis.Options{
    Addr:     "redis:6379",
    Password: "",
    DB:       0,
  })

  if pong, err := client.Ping().Result(); err {
    fmt.Println("Result: ", pong)
  } else {
    fmt.Println("Error (",err,")")
  }
}
