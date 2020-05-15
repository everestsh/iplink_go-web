
//go语言怎么调用shell脚本文件
package main
import (
  "fmt"
  "os/exec"
)

func main(){
  //cmd := exec.Command("ls", "-l", "-a")
  cmd := exec.Command("/bin/bash", "./1.sh")
  bytes, err := cmd.Output()
  if err != nil {
    fmt.Println("cmd.Output:", err)
    return
  }
  fmt.Println(string(bytes))
}
