//
//终端标准输出tmp.txt
package main

import(
  "fmt"
  "os"
  "os/exec"
)

func main()  {
  cmd := exec.Command("cat")
  stdin, err := cmd.StdinPipe()
  if err != nil {
    fmt.Println(err)
  }

  _, err = stdin.Write([]byte("tmp.txt"))
  if err != nil {
    fmt.Println(err)
  }
  stdin.Close()
  cmd.Stdout = os.Stdout //终端标准输出tmp.txt
  cmd.Start()
}
