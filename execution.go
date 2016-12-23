package main

import (
	"fmt"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func executeTestSuite() {
	for _, cmd := range ConfigFile.Command {
		executeTest(cmd)
	}
}

func executeTest(cmd string) {
	containerOpts := docker.CreateContainerOptions{
		// Name: "mytestcontainer" + cmd,
		Config: &docker.Config{
			Image:        ConfigFile.Image,
			Cmd:          []string{cmd},
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

	err = client.AttachToContainer(attachOpts)

	if err != nil {
		panic(fmt.Errorf("failed to attach container: %v", err))
	}

	// wait on the container
	exitCode, _ := client.WaitContainer(container.ID)
	fmt.Printf("EXIT CODE: %d", exitCode)

	// remove the container
	removeOpts := docker.RemoveContainerOptions{
		ID:            container.ID,
		Force:         true,
		RemoveVolumes: true,
	}
	client.RemoveContainer(removeOpts)
}
