package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

type dockerClient struct {
	*docker.Client
}

func (c *dockerClient) initialize() {
	fmt.Println("Connecting to Docker daemon...")

	var err error
	c.Client, err = docker.NewClient("unix:///var/run/docker.sock")

	if err != nil {
		panic(fmt.Errorf("error connecting to daemon: %s", err))
	}

	err = client.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging daemon: %s", err))
	}
}

func (c *dockerClient) isConnected() error {
	err := c.Ping()

	if err != nil {
		return err
	}

	return nil
}
