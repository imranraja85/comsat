package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

type dockerClient struct {
	*docker.Client
}

func (c *dockerClient) initialize() {
	var err error
	c.Client, err = docker.NewClient("unix:///var/run/docker.sock")

	if err != nil {
		panic(fmt.Errorf("error connecting to daemon: %s", err))
	}
}
