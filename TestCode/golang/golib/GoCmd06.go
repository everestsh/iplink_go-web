
//echo "hello world" | tr 'a-z' 'A-Z'
//输出结果是：HELLO WORLD
package main

import (
  "fmt"
  "os/exec"
  "log"
  "strings"
  "bytes"
)

func main()  {
  cmd := exec.Command("tr", "a-z", "A-Z")
  cmd.Stdin = strings.NewReader("some input")
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("in all caps: %q\n", out.String()) //in all caps: " SOME INPUT"
}
