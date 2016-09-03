package docktor

import (
	"fmt"

	dockerClient "github.com/fsouza/go-dockerclient"
)

const (
	dockerEndpoint = "unix:///var/run/docker.sock"
)

func main() {
	connect(Docker{path: dockerEndpoint})
}

func connect(options Docker) {
	client, err := dockerClient.NewClient(options.path)
	if err != nil {
		panic("Unable to connect to Docker")
	}
	images, _ := client.ListImages(dockerClient.ListImagesOptions{})

	for _, image := range images {
		fmt.Printf(image.ID + "\n")
	}
}

// Docker options for connecting to
type Docker struct {
	path string // unix socket path. defaults to unix:///var/run/docker.sock

}
