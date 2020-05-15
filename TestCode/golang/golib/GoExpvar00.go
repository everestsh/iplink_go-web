

package main

import (
  "expvar"
  "net/http"
  "fmt"
)

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request)  {
  visits.Add(1)
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main()  {
  https.HandlerFunc("/", handler)
  http.ListenAndServer(":8080", nil)
}
