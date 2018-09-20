package main

import (
  "net/http"
  "net/url"
  "fmt"
  "log"
  "strconv"
  "os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  u, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
  prvi, err := strconv.Atoi(q.Get("a"))
    	if err != nil {
        	fmt.Println(err)
        	os.Exit(2)
    		}
  fmt.Println(prvi)

	drugi, err := strconv.Atoi(q.Get("b"))
    	if err != nil {
        	fmt.Println(err)
        	os.Exit(2)
    		}
  
  test := strconv.Itoa(prvi + drugi)

  w.Write([]byte("jel radi?" + test))
}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8081", nil); err != nil {
    panic(err)
  }
}
