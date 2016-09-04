package api

import (
	docker "github.com/fsouza/go-dockerclient"
)

// Get Host images and containers
func Get(client *docker.Client) Host {
	containers := GetContainers(client)
	images := GetImages(client)

	host := Host{containers, images}
	return host
}

// GetImages List host images
func GetImages(client *docker.Client) []docker.APIImages {
	images, err := client.ListImages(docker.ListImagesOptions{})
	if err != nil {
		panic(err.Error())
	}
	return images
}

// GetContainers List host containers
func GetContainers(client *docker.Client) []docker.APIContainers {
	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		panic(err.Error())
	}
	return containers
}

// GetClient Return a Docker client connected to Host
func GetClient(hostname string) *docker.Client {
	client, error := docker.NewClient(hostname)
	if error != nil {
		panic("Unable to connect to Docker Host '" + hostname + "'")
	}
	return client
}

// Host ...
type Host struct {
	containers []docker.APIContainers
	images     []docker.APIImages
}
