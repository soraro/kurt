package cmd

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func collect(clientset *kubernetes.Clientset, namespace []string, labels []string) {

	namespaceTracker = make(map[string]int32)
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
				trackPods(v.ObjectMeta.Name, restarts)
				trackNamespaces(v.ObjectMeta.Namespace, restarts)
				trackLabels(labels, v.ObjectMeta.Labels, restarts)
			}
		}
	}
	showResults()

}

func trackNamespaces(namespace string, count int32) {
	namespaceTracker[namespace] += count
}

func trackPods(pod string, count int32) {
	podTracker[pod] = count
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
