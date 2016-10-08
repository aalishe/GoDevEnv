package main

/* Example #38: https://gobyexample.com/stateful-goroutines */

import (
  "fmt"
  "math/rand"
  "sync/atomic"
  "time"
)

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {
  var opsR int64 = 0
  var opsW int64 = 0

  reads := make(chan *readOp)
  writes:= make(chan *writeOp)

  go func() {
    var state = make(map[int]int)

    for {
      select {
      case read := <-reads:
        read.resp <- state[read.key]
      case write := <-writes:
        state[write.key] = write.val
        write.resp <- true
      }
    }
  }()

  for r := 0; r < 100; r++ {
    go func() {
      for {
        read := &readOp{
          key: rand.Intn(5),
          resp: make(chan int),
        }
        reads <- read
        <-read.resp
        atomic.AddInt64(&opsR, 1)
      }
    }()
  }

  for w := 0; w < 10; w++ {
    go func() {
      for {
        write := &writeOp{
          key: rand.Intn(5),
          val: rand.Intn(100),
          resp: make(chan bool),
        }
        writes <- write
        <-write.resp
        atomic.AddInt64(&opsW, 1)
      }
    }()
  }

  time.Sleep(time.Second)

  opsRFinal := atomic.LoadInt64(&opsR)
  opsWFinal := atomic.LoadInt64(&opsW)
  fmt.Println("Read Ops:", opsRFinal)
  fmt.Println("Write Ops:", opsWFinal)
}
