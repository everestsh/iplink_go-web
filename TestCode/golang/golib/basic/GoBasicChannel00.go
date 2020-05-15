package main

import "fmt"

func main() {
	fmt.Println("function begins ...")
	c := make(chan bool)
	go func() {
		fmt.Println("fun has benn called.")
		close(c)
	}()
	<-c
	fmt.Println("Completed.")
}
