package cmd

import (
	"reflect"
	"testing"
)

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
