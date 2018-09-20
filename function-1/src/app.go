package main

import (
  "net/http"
  "strconv"
  "os"
  "fmt"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  x := r.Form.Get("a")
  y := r.Form.Get("b")

  prvi, err := strconv.Atoi(x)
    	if err != nil {
        	fmt.Println(err)
        	os.Exit(2)
    		}

	drugi, err := strconv.Atoi(y)
    	if err != nil {
        	fmt.Println(err)
        	os.Exit(2)
    		}

  test := strconv.Itoa(prvi + drugi)

  w.Write([]byte("Sum of values is: " + test))

}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8081", nil); err != nil {
    panic(err)
  }
}
