package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	args, n := getArgsP()
	if n != 0 {
		fmt.Println("n = ", n)
	} else {
		fmt.Println("n = ", n)
		fmt.Println("args = ", args)
	}
	/*
		var s = "-MemTotal: 1001332 kB"
		//var valid = regexp.MustCompile("[0-9]")
		//var valid = regexp.MustCompile("^-")
		var valid, _ = regexp.Compile("^-")
		on := valid.Find([]byte(s))
		fmt.Println("Find:", string(on))
		fmt.Println(valid.FindAllStringSubmatch(s, -1))
		//fmt.Println(GetFields(s))

		a := "I am learning Go language"

		re, _ := regexp.Compile("[a-z]{2,4}")

		//查找符合正则的第一个
		one := re.Find([]byte(a))
		fmt.Println("Find:", string(one))
	*/
}

//getArgsP 获取参数列表的带有-的
func getArgsP() (string, int) {
	var v, _ = regexp.Compile("^-")
	//b := false
	i := 0
	for i, n := range os.Args[0:] {
		o := v.Find([]byte(os.Args[i]))
		if string(o) == "-" {
			//b = true
			fmt.Println(n)
			fmt.Println(i)
			//fmt.Println(b)
			return n, i
		}
	}
	return "", i
}

/*
func GetFields(str string) []string {
	var sub = strings.Fields(str)
	return sub
}
*/
