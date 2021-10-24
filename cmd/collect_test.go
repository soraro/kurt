package cmd

import (
	"testing"
)

func TestTrackNamespaces(t *testing.T) {
	namespaceTracker = make(map[string]int32)
	trackNamespaces("ns1", 5)
	trackNamespaces("ns1", 2)

	if namespaceTracker["ns1"] != int32(7) {
		t.Errorf("ns1 namespace expected to have a count of 7 but instead shows: %v", namespaceTracker["ns1"])
	}
}

func TestTrackNodes(t *testing.T) {
	nodeTracker = make(map[string]int32)
	trackNodes("node01", 5)
	trackNodes("node01", 2)
	trackNodes("node02", 2)

	if nodeTracker["node01"] != int32(7) {
		t.Errorf("node01 node expected to have a count of 7 but instead shows: %v", nodeTracker["node01"])
	}
}

func TestTrackPods(t *testing.T) {
	podTracker = make(map[string]int32)
	// Test that a pod with the same name in a different namespace is held uniquely in the map
	trackPods("pod1", "default", 3)
	trackPods("pod1", "other", 2)
	if podTracker["default:pod1"] != 3 {
		t.Errorf("pod1 pod expected to have a count of 3 but instead shows: %v", podTracker["pod1"])
	}
}

func TestTrackLabels(t *testing.T) {
	labelTracker = make(map[string]int32)
	tlabels := []string{"app", "k8s-app"}
	plabelsA := map[string]string{
		"app":   "app1",
		"other": "label",
	}
	plabelsB := map[string]string{
		"k8s-app": "app2",
	}

	trackLabels(tlabels, plabelsA, 3)
	trackLabels(tlabels, plabelsB, 5)

	// other:label should not exist because it is not defined in tlabels
	if labelTracker["other:label"] != int32(0) {
		t.Errorf("other:label should not exist because it was not defined in user-defined tlabels")
	}

	if labelTracker["app:app1"] != int32(3) {
		t.Errorf("app:app1 should be equal to 3 since it is defined in the tlabels but instead shows: %v", labelTracker["app:app1"])
	}

	if labelTracker["k8s-app:app2"] != int32(5) {
		t.Errorf("ks-app:app2 should be equal to 5 since it is defined in the tlabels but instead shows: %v", labelTracker["k8s-app:app2"])
	}
}
