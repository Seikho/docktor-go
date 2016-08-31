package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	client := docker.StartExecOptions{
		Detach:      true,
		Tty:         true,
		RawTerminal: true,
	}
	fmt.Println(client.RawTerminal)
}
