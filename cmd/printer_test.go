package cmd

import (
	"reflect"
	"testing"
)

func Test_returnSortedLimit(t *testing.T) {
	type args struct {
		data    map[string]int32
		limit   int
		parseNS bool
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
				limit:   5,
				parseNS: false,
			},
			want: ItemList{
				Item{
					Name:      "test4",
					Count:     9,
					Namespace: "",
				},
				Item{
					Name:      "test3",
					Count:     8,
					Namespace: "",
				},
				Item{
					Name:      "test2",
					Count:     7,
					Namespace: "",
				},
				Item{
					Name:      "test1",
					Count:     5,
					Namespace: "",
				},
				Item{
					Name:      "test6",
					Count:     4,
					Namespace: "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnSortedLimit(tt.args.data, tt.args.limit, tt.args.parseNS); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnSortedLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
