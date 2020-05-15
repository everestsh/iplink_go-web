
package main

import (
  "fmt"
  "os/exec"
)

func main ()  {
  //cmd := exec.Command("ls", "-l", "-a")
  //cmd := exec.Command("/bin/sh", "-c", "ls")
  cmd := exec.Command("ls")
  out, err := cmd.Output()
  if err != nil{
    //log.Fatal(err)
    fmt.Println(err)
  }
  fmt.Println(string(out))
}
