package main

import (
  "net/http"
  "fmt"
  "net"
)

func returnIp(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  url := r.Form.Get("url")
  ip, err := net.LookupIP(url)
  if err != nil {
      fmt.Println(err)
    }
  ipString := ip[0].String()
  w.Write([]byte("IP address of \"" + url + "\" is: " + ipString))
}

func main() {
  http.HandleFunc("/", returnIp)
  if err := http.ListenAndServe(":8082", nil); err != nil {
    panic(err)
  }
}
