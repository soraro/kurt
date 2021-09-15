package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

// Handle setting up cluster auth and return clientset
func auth() *kubernetes.Clientset {

	kubeconfig := getConfigPath()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset

}

func getConfigPath() string {

	if os.Getenv("KUBECONFIG") == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		return filepath.Join(homedir, ".kube", "config")
	} else {
		return os.Getenv("KUBECONFIG")
	}

}
