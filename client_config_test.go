package main

import "testing"

// this does a live test. replace with something that stubs a success/failure http call to the docker remote api
func TestIsClientConnected(t *testing.T) {
	client.initialize()
	err := client.isConnected()

	if err != nil {
		t.Fatal("Expected nil error, got error: %v", err)
	}
}
