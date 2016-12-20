package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

var client *docker.Client

func initializeDockerClient() {
	var err error
	client, err = docker.NewClient("unix:///var/run/docker.sock")

	if err != nil {
		panic(fmt.Errorf("error connecting to daemon: %s", err))
	}
}
