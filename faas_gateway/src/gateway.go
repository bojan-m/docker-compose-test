package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "net/url"
)

func readUrl(w http.ResponseWriter, r *http.Request) {
  u, err := url.Parse(r.URL.Path)
	if err != nil {
		fmt.Println(err)
	}

  q := r.URL.Query()

  switch {
    case u.Path == "/function-1" && len(q) > 1:
      response, err := http.PostForm("http://function-1:8081/",
            url.Values{"a" : q["a"], "b" : q["b"]})
      if err != nil {
          fmt.Println(err)
      }
      defer response.Body.Close()
      test, err := ioutil.ReadAll(response.Body)
      if err != nil {
        fmt.Println(err)
      }
      stringBody := string(test)
      w.Write([]byte(stringBody))

    case u.Path == "/function-2" && len(q) == 1:
      response, err := http.PostForm("http://function-2:8082/",
            url.Values{ "url" : q["url"]})
      if err != nil {
          fmt.Println(err)
      }
      defer response.Body.Close()
      test, err := ioutil.ReadAll(response.Body)
      if err != nil {
        fmt.Println(err)
      }
      stringBody := string(test)
      w.Write([]byte(stringBody))

    default:
      message := "Invalid request"
      w.Write([]byte(message))
    }
  }

func main() {
  http.HandleFunc("/", readUrl)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
