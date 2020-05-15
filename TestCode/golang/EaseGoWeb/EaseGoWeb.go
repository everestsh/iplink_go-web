package main

import (
	"net/http"
	"os"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir("/Users/raylea/Workspace/github/iplinkme_go-web/TestCode/golang/EaseGoWeb/microembedded.com")))
	//http.Handle("./", http.FileServer(http.Dir("./microembedded.com")))
	//http.ListenAndServe(":8080", nil)
	//2
	//http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))

	//3
	http.ListenAndServe(":8080", http.FileServer(http.Dir(os.Args[1])))
}
