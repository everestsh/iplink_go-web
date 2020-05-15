//
//The output is: widuu test!
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	//  "io/ioutil"
)

func main() {
	var output bytes.Buffer
	cmd := exec.Command("cat")
	cmd.Stdout = &output
	stdin, _ := cmd.StdinPipe() //指向cmd命令的stdout
	cmd.Start()
	stdin.Write([]byte("widuu test"))
	stdin.Close()
	cmd.Wait()
	fmt.Printf("The Output is: %s\n", output.Bytes())
}
