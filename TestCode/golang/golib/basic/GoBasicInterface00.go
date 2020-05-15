package main

import "fmt"

type person struct {
	id      int
	name    string
	country string
}

type interface_preson interface {
	introduction()
}

func (this *person) introduction() {
	fmt.Println("My name is : this.name")
}

func main() {
	var liumiao person
	liumiao.id = 1001
	liumiao.name = "liumiao"
	liumiao.country = "China"

	fmt.Println("liumaio=", liumiao)
	liumiao.introduction()
}
