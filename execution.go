package main

import (
	"fmt"
	"os"
	"sync"

	docker "github.com/fsouza/go-dockerclient"
)

type Execution struct {
	Cmd string
}

// executeTestSuite iterates over each command in the config file and
// executes it in a container
func executeTestSuite() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, cmd := range ConfigFile.Command {
		e := Execution{cmd}
		wg.Add(1)
		go e.executeTest(ch, &wg)
	}

	// decide what to do here when a container has an
	// exit code of something other than 0. Do we want to:
	//   1) Ignore it and continue exiting other commands
	//   2) Stop and remove all containers and report something 'failed'
	go func() {
		for {
			exitCode := <-ch
			fmt.Println(exitCode)
			//if exitCode != 0 {
			//	os.Exit(exitCode)
			//}
		}
	}()

	wg.Wait()
}

// executeTest executes a command by creating a container, running the container,
// and then redirect stdout to the process.
func (e *Execution) executeTest(c chan int, wg *sync.WaitGroup) {
	fmt.Println("Starting Execution of:", e.Cmd)

	containerOpts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:        ConfigFile.Image,
			Cmd:          []string{e.Cmd},
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
	c <- exitCode

	// remove the container
	removeOpts := docker.RemoveContainerOptions{
		ID:            container.ID,
		Force:         true,
		RemoveVolumes: true,
	}
	client.RemoveContainer(removeOpts)

	wg.Done()
}
