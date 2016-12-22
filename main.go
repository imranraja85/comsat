package main

func init() {
	initializeDockerClient()
	initializeStackConfig()
}

func main() {
	executeTestSuite()
}
