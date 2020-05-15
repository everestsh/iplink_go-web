
//让Go语言调用外部程序时支持管道符
//ls | wc -l
package main

import (
  "os"
  "os/exec"
)

func main ()  {
  run2()
}


func run2()  {
  c1 := exec.Command("ls")
//  c2 := exec.Command("wc", "-l")
//  c2.Stdin, _ = c1.StdoutPipe()
  c1.Stdout = os.Stdout
  c1.Stderr = os.Stderr
  c1.Start()
  c1.Run()
//  c2.Wait()

}
