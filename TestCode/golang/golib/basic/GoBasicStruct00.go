package main

import "fmt"

func main() {
	type person struct {
		id      int
		name    string
		country string
	}
	//def and init 1
	var liumiao person
	liumiao.id = 1001
	liumiao.name = "liumiao"
	liumiao.country = "China"

	fmt.Println("liumiao=", liumiao)

	//def and init 2
	michael := person{1002, "micheal", "PRC"}
	fmt.Println("micheal=", michael)
	fmt.Printf("micheal= %#v\n", michael)
}
