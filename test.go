// package main
//
// import "fmt"
//
// func main() {
// 	fmt.Printf("hello, world\n")
// }

package main

import (
  "net/http"
  "strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  if message == "function-1" {
    http.Redirect(w, r, "http://google.com", 303) 
  } else {
    message = "zasto ne radi:" + message
    w.Write([]byte(message))
  }
}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
