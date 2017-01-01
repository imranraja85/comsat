package main

var client dockerClient

func init() {
	client.initialize()
	initializeStackConfig()
}

func main() {
	executeTestSuite()
}
