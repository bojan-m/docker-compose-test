package main

import (
  	"context"
    "fmt"
    "net/http"
    "io/ioutil"
    "net/url"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func getFunctions() map[string][]string {
  cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		fmt.Println(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Println(err)
	}

  s := make(map[string][]string)

	for _, container := range containers {
    s[container.Labels["function.name"]] = []string{container.Labels["function.port"], container.Labels["function.hostname"]}
	}

  return s
}

func readUrl(w http.ResponseWriter, r *http.Request) {
  u, err := url.Parse(r.URL.Path)
	if err != nil {
		fmt.Println(err)
	}

  s := getFunctions()

  q := r.URL.Query()

  value, ok := s[u.Path]
  if ok && len(q) > 0 {
    response, err := http.PostForm("http://"+value[1]+":"+value[0]+"/", q)
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
    } else {
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
