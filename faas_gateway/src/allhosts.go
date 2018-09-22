package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//for testing the API

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

  s := make(map[string][]string)

	for _, container := range containers {
    s[container.Labels["function.name"]] = []string{container.Labels["function.port"], container.Labels["function.hostname"]}
		fmt.Printf("%s %s %s\n", container.Labels["function.name"], container.Labels["function.port"], container.Labels["function.hostname"])
	}

  fmt.Println(s)


}
