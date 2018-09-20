package main

import (
  // "fmt"
  // "os"
  "net/http"
  "strings"
  "log"
  "io/ioutil"
)

func readUrl(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  params := strings.TrimSuffix(message, "?")
  message = strings.TrimPrefix(message, "/")
  switch {
    case message == "function-1":
      response, err := http.Get("http://function-1:8081/" + string(params))
      if err != nil {
          log.Fatal(err)
      }
      defer response.Body.Close()
      test, err := ioutil.ReadAll(response.Body)
      if err != nil {
        log.Fatal(err)
      }
      stringBody := string(test)
      w.Write([]byte(stringBody))

    case message == "function-2":
      response, err := http.Get("http://function-2:8082")
      if err != nil {
          log.Fatal(err)
      }
      defer response.Body.Close()
      test, err := ioutil.ReadAll(response.Body)
      if err != nil {
        log.Fatal(err)
      }
      stringBody := string(test)
      w.Write([]byte(stringBody))

    default:
      message = "wrong request"
      w.Write([]byte(message))
    }
  }

func main() {
  http.HandleFunc("/", readUrl)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
