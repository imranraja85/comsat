package main

var client dockerClient

func main() {
	client.initialize()
	initializeStackConfig()
	executeTestSuite()
}
