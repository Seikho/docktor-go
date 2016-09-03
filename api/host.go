package api

import (
	docker "github.com/fsouza/go-dockerclient"
)

func Get(client *docker.Client) Host {
	containers := GetContainers(client)
	images := GetImages(client)

	host := Host{containers, images}
	return host
}

func GetImages(client *docker.Client) []docker.APIImages {
	images, err := client.ListImages(docker.ListImagesOptions{})
	if err != nil {
		panic(err.Error())
	}
	return images
}

func GetContainers(client *docker.Client) []docker.APIContainers {
	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		panic(err.Error())
	}
	return containers
}

// Host ...
type Host struct {
	containers []docker.APIContainers
	images     []docker.APIImages
}
