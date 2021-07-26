package collect

import "testing"

func TestCollect(t *testing.T) {
	type args struct {
		namespace string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test1",
			args: args{namespace: "ns1"},
			want: "ns1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.args.namespace); got != tt.want {
				t.Errorf("Collect() = %v, want %v", got, tt.want)
			}
		})
	}
}
