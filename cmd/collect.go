package cmd

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func collect(clientset *kubernetes.Clientset, namespace []string, labels []string) {

	if limitFlag < 0 {
		log.Fatal("FATAL CONFIGURATION: --limit flag value must not be negative.")
	}

	if !(output == "standard" || output == "yaml" || output == "json") {
		log.Fatal("FATAL CONFIGURATION: --output flag can only be: standard, json, yaml")
	}

	namespaceTracker = make(map[string]int32)
	nodeTracker = make(map[string]int32)
	podTracker = make(map[string]int32)
	labelTracker = make(map[string]int32)

	for _, ns := range namespace {
		for _, lb := range labels {
			pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{LabelSelector: lb})
			if err != nil {
				log.Fatal(err.Error())
			}

			for _, v := range pods.Items {
				restarts := int32(0)
				for _, vv := range v.Status.ContainerStatuses {
					restarts += vv.RestartCount
				}
				trackPods(v.ObjectMeta.Name, v.ObjectMeta.Namespace, restarts)
				trackNamespaces(v.ObjectMeta.Namespace, restarts)
				trackLabels(labels, v.ObjectMeta.Labels, restarts)
				trackNodes(v.Spec.NodeName, restarts)
			}
		}
	}
	showResults()

}

func trackNamespaces(namespace string, count int32) {
	namespaceTracker[namespace] += count
}

func trackNodes(node string, count int32) {
	nodeTracker[node] += count
}

func trackPods(pod, namespace string, count int32) {
	podTracker[namespace+":"+pod] = count
}

// plabels = Pod Labels
// tlabels = (User-defined) tracking labels
func trackLabels(tlabels []string, plabels map[string]string, count int32) {
	// range through all the labels specified in the -l CLI flag
	for _, l := range tlabels {
		// range through plabels to see if any match the user specified labels. If so, add it to the map
		// the default value "*" will match everything
		for k, v := range plabels {
			if l == k || l == "" {
				labelTracker[k+":"+v] += count
			}
		}
	}
}
