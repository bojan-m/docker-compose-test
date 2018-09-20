package main

import (
  "net/http"
  "strings"
  "fmt"
  "net"
)

func returnIp(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "hello function 2" + message
  test := net.LookupIP("www.google.com")


  fmt.Println(r)
  fmt.Println(test)

  w.Write([]byte(message))
  w.Write([]byte(test))
}

func main() {
  http.HandleFunc("/", returnIp)
  if err := http.ListenAndServe(":8082", nil); err != nil {
    panic(err)
  }
}
