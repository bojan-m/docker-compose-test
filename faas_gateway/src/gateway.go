package main

import (
    // "fmt"
  // "os"
    "net/http"
    // "strings"
    "log"
    "io/ioutil"
    "net/url"
)

func readUrl(w http.ResponseWriter, r *http.Request) {
  u, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

  q := r.URL.Query()

  switch {
    case u.Path == "/function-1" && len(q) > 1:
      response, err := http.PostForm("http://function-1:8081/",
            url.Values{"a" : q["a"], "b" : q["b"]})
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

    case u.Path == "/function-2":
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
      message := "wrong request"
      w.Write([]byte(message))
    }
  }

func main() {
  http.HandleFunc("/", readUrl)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
