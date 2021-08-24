package main

import (
	"os"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	// make sure that if a KUBECONFIG environment variable is set, the result returns that specified path
	os.Setenv("KUBECONFIG", "/my/path/config")
	path := getConfigPath()

	if path != "/my/path/config" {
		t.Errorf("KUBECONFIG env var is not being respected")
	}
	os.Unsetenv("KUBECONFIG")
}
