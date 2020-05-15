package main

import "fmt"

func main() {
	m := make(map[string]int) //使用make创建一个空的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)      //输出map[one:1 two:2 three:3]
	fmt.Println(len(m)) //输出3

	v := m["two"]  //从map里取值
	fmt.Println(v) //输出2

}
