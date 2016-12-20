package main

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
)

func init() {
	initializeDockerClient()
	initializeStackConfig()
	executeCommand()
}

func main() {
}

func executeCommand() {
	// loop over the commands and execute each command in its own container
	// set up container options
	containerOpts := docker.CreateContainerOptions{
		Name: "mytestcontainer",
		Config: &docker.Config{
			Image:        ConfigFile.Image,
			Cmd:          ConfigFile.Command,
			AttachStdout: true,
			AttachStderr: true,
		},
	}

	// create the container
	container, err := client.CreateContainer(containerOpts)
	if err != nil {
		panic(fmt.Errorf("error: %v", err))
	}

	// run the container
	client.StartContainer(container.ID, nil)

	// attach the container to stdout of client
	attachOpts := docker.AttachToContainerOptions{
		Container:    container.ID,
		Stdin:        true,
		Stdout:       true,
		Stderr:       true,
		Stream:       true,
		Logs:         true,
		OutputStream: os.Stdout,
	}

	_ = client.AttachToContainer(attachOpts)
	if err != nil {
		panic(fmt.Errorf("failed to attach container: %v", err))
	}

	// remove the container
	removeOpts := docker.RemoveContainerOptions{
		ID:            container.ID,
		Force:         true,
		RemoveVolumes: true,
	}
	client.RemoveContainer(removeOpts)
}
