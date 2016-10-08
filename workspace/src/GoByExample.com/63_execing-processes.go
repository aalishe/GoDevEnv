package main

/* Example #63: https://gobyexample.com/execing-processes */

import "fmt"
import "syscall"
import "os"
import "os/exec"

func main() {
  binary, lookErr := exec.LookPath("ls")
  if lookErr != nil {
    panic(lookErr)
  }
  fmt.Println(string(binary))

  args := []string{"ls", "-a", "-l", "-h", ".."}

  env := os.Environ()

  execErr := syscall.Exec(binary, args, env)
  // Not at this point unless error
  fmt.Println(args)
  if execErr != nil {
    panic(execErr)
  }
}
