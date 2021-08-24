package main

import (
	"flag"
	"fmt"
)

type labels []string

func (x *labels) Set(value string) error {
	*x = append(*x, value)
	return nil
}

func (x *labels) String() string {
	return fmt.Sprint(*x)
}

func main() {
	var namespace string
	var labels labels

	flag.Usage = func() {
		fmt.Println("kurt: KUbernetes Reboot Tracker")
		flag.PrintDefaults()
	}

	flag.StringVar(&namespace, "n", "", "Specify namespace for kurt to collect reboot metrics.\nLeave blank to collect in all namespaces.")
	flag.Var(&labels, "l", "Specify multiple times for the label keys you want to see.\nFor example: -l app")
	flag.Parse()

	clientset := auth()
	collect(clientset, namespace, labels)

}
