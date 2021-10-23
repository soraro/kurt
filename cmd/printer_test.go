package cmd

import (
	"reflect"
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

func Test_returnSortedLimit(t *testing.T) {
	type args struct {
		data  map[string]int32
		limit int
	}
	tests := []struct {
		name string
		args args
		want PairList
	}{
		{
			name: "test1",
			args: args{
				data: map[string]int32{
					"test1": 5,
					"test2": 7,
					"test3": 8,
					"test4": 9,
					"test5": 0,
					"test6": 4,
				},
				limit: 5,
			},
			want: PairList{
				Pair{
					Key:   "test4",
					Value: 9,
				},
				Pair{
					Key:   "test3",
					Value: 8,
				},
				Pair{
					Key:   "test2",
					Value: 7,
				},
				Pair{
					Key:   "test1",
					Value: 5,
				},
				Pair{
					Key:   "test6",
					Value: 4,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnSortedLimit(tt.args.data, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnSortedLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
