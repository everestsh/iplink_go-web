//http://127.0.0.1:9090/search?q=dotnet
/*
map[q:[dotnet]]
path: /search
scheme:
[]
key: q
val: dotnet
*/
//http://localhost:9090/?url_long=111&url_long=222
/*
map[url_long:[111 222]]
path: /
scheme:
[111 222]
key: url_long
val: 111222
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                        //解析参数，默认是不会解析的
	fmt.Println(r.Form)                  //这些信息是输出到服务器端的打印信息  << map[q:[dotnet]] >>
	fmt.Println("path:", r.URL.Path)     //  << /search >>
	fmt.Println("scheme:", r.URL.Scheme) // << >>
	fmt.Println(r.Form["url_long"])      //  []
	for k, v := range r.Form {
		fmt.Println("key:", k)                     //  key: q
		fmt.Println("val:", strings.Join((v), "")) //  val: dotnet
	}
	//fmt.Println(w, "hello world")
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	//http.ListenAndServe(":9090", nil)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
