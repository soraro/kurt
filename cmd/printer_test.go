package cmd

import (
	"reflect"
	"testing"
)

func Test_returnSortedLimit(t *testing.T) {
	type args struct {
		data       map[string]int32
		limit      int
		parseNS    bool
		containers map[string]map[string]int32
	}
	tests := []struct {
		name string
		args args
		want ItemList
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
				limit:      5,
				parseNS:    false,
				containers: nil,
			},
			want: ItemList{
				Item{
					Name:       "test4",
					Count:      9,
					Namespace:  "",
					Containers: nil,
				},
				Item{
					Name:       "test3",
					Count:      8,
					Namespace:  "",
					Containers: nil,
				},
				Item{
					Name:       "test2",
					Count:      7,
					Namespace:  "",
					Containers: nil,
				},
				Item{
					Name:       "test1",
					Count:      5,
					Namespace:  "",
					Containers: nil,
				},
				Item{
					Name:       "test6",
					Count:      4,
					Namespace:  "",
					Containers: nil,
				},
			},
		},
		{
			name: "test2",
			args: args{
				data: map[string]int32{
					"test1:pod1": 5,
					"test2:pod2": 7,
					"test2:pod3": 8,
					"test1:pod4": 9,
					"test2:pod5": 0,
					"test1:pod6": 4,
				},
				limit:   5,
				parseNS: true,
				containers: map[string]map[string]int32{
					"test1:pod1": {"container1": 5},
					"test2:pod2": {"container1": 7},
					"test2:pod3": {"container1": 8},
					"test1:pod4": {"container3": 2, "container1": 4, "container2": 3},
					"test2:pod5": {"container1": 0},
					"test1:pod6": {"container1": 4},
				},
			},
			want: ItemList{
				Item{
					Name:       "pod4",
					Count:      9,
					Namespace:  "test1",
					Containers: Containers{Container{"container1", 4}, Container{"container2", 3}, Container{"container3", 2}},
				},
				Item{
					Name:       "pod3",
					Count:      8,
					Namespace:  "test2",
					Containers: Containers{Container{"container1", 8}},
				},
				Item{
					Name:       "pod2",
					Count:      7,
					Namespace:  "test2",
					Containers: Containers{Container{"container1", 7}},
				},
				Item{
					Name:       "pod1",
					Count:      5,
					Namespace:  "test1",
					Containers: Containers{Container{"container1", 5}},
				},
				Item{
					Name:       "pod6",
					Count:      4,
					Namespace:  "test1",
					Containers: Containers{Container{"container1", 4}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnSortedLimit(tt.args.data, tt.args.limit, tt.args.parseNS, tt.args.containers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnSortedLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
