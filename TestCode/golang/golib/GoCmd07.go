
//ls
//
package main

import (
  "fmt"
  "os/exec"
//  "strings"
)

func main()  {
  cmd := exec.Command("ls")
  out, err := cmd.Output()
  if err != nil{
    fmt.Println("err")
  }
  fmt.Println(string(out))
}
