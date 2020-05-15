

//
package main

import (
  "io"
  "log"
  "net/http"
)

func HelloGoServer(w http.ResponseWriter, r *http.Request)  {
  io.WriteString(w, "Hello, this is a GoServer")
}

func main()  {
  http.HandleFunc("/", HelloGoServer)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServer", err)
  }
}
