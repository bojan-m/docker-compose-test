package main

import (
  "net/http"
  "strconv"
  "fmt"
)

func sumOfValues(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  x := r.Form.Get("a")
  y := r.Form.Get("b")

  prvi, err := strconv.Atoi(x)
    	if err != nil {
        	fmt.Println(err)
    		}

	drugi, err := strconv.Atoi(y)
    	if err != nil {
        	fmt.Println(err)
    		}

  test := strconv.Itoa(prvi + drugi)

  w.Write([]byte("Sum of values is: " + test))

}

func main() {
  http.HandleFunc("/", sumOfValues)
  if err := http.ListenAndServe(":8081", nil); err != nil {
    panic(err)
  }
}
