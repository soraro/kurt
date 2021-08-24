package main

import (
	"testing"
)

func TestReturnSorted(t *testing.T) {
	podMap := map[string]int32{
		"pod1": int32(2),
		"pod2": int32(5),
		"pod3": int32(0),
		"pod4": int32(2),
	}
	pl := returnSorted(podMap)

	// Check that pod2 is the first index since it has the highest restart count
	if pl[0].Key != "pod2" {
		t.Errorf("pod2 should be the first index, but instead got: %v", pl[0].Key)
	}

	// Check that pod3 is the last index since it has the lowest restart count
	if pl[len(pl)-1].Key != "pod3" {
		t.Errorf("pod3 should be the last index, but instead got: %v", pl[len(pl)-1].Key)
	}
}
