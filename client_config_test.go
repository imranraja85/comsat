package main

import "testing"

func TestConfigureClient(t *testing.T) {
	client.initialize() // connects and checks for a docker client connection

	hello := "world"

	if hello != "world" {
		t.Errorf("Expected %s, got %s", "world", "worl")
	}

}
