package main

import (
	"fmt"
	"os"
)

func main() {
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}
	s := "hello,world"
	s1 := "$hello,$world $finish"
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))
}
