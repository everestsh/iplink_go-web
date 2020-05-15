
package main

import (
  "fmt"
  "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "hello World")
}

func main() {
  http.HandleFunc("/", IndexHandler)
  http.ListenAndServe(":8000", nil)
}
