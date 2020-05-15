package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir(
			"/Users/raylea/Workspace/github/iplinkme_go-web/iplinkme_go-web/TestCode/jQuery/DOMmanipulation/")))
	//http.Handle("./", http.FileServer(http.Dir("./microembedded.com")))
	http.ListenAndServe(":8080", nil)
}
