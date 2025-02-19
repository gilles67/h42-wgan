package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func main() {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	networks, err := cli.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%-30s %s\n", ctr.Names[0], ctr.Labels)
	}

	for _, net := range networks {
		fmt.Printf("%-30s %s\n", net.Name, net.Labels)
	}
}
