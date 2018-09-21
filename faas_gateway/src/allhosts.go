package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

  s := make([]string, 0)

	for _, container := range containers {
    s = append(s, container.Labels["function.name"], container.Labels["function.port"], container.Labels["function.hostname"])
		fmt.Printf("%s %s %s\n", container.Labels["function.name"], container.Labels["function.port"], container.Labels["function.hostname"])
	}

  fmt.Println(s)


}
