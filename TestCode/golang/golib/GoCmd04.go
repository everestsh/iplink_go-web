
//ps -eaf|grep "nginx: master"|grep -v "grep"|awk '{print $2}'

package main

import (
  "os"
  "os/exec"
)

func main()  {
//  run1()
  run2()
  run3()
}
/*
func run1()  {
  cmd := exec.Command("ls", "|", "wc", "-l")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Start()
  cmd.Run()
  cmd.Wait()
}
*/

func run2()  {
  c1 := exec.Command("ls")
  c2 := exec.Command("wc", "-l")
  c2.Stdin, _ = c1.StdoutPipe()
  c2.Stdout = os.Stdout
  c2.Stderr = os.Stderr
  c2.Start()
  c1.Run()
  c2.Wait()
}

func run3()  {
  c1 := exec.Command("ps", "-eaf")
  c2 := exec.Command("grep", `"nginx: master"`)
  c3 := exec.Command("grep", "-v", `"grep"`)
  c4 := exec.Command("awk", `"{print $2}"`)
  c2.Stdin, _ = c1.StdoutPipe()
  c3.Stdin, _ = c2.StdoutPipe()
  c4.Stdin, _ = c3.StdoutPipe()

  c4.Stdout = os.Stdout
  c4.Stderr = os.Stderr

  c4.Start()
  c3.Start()
  c2.Start()
  c1.Run()
  c4.Wait()
}
