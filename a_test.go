package codecov_gh_actions

import "testing"

func Test_a(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{i: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := a(tt.args.i); got != tt.want {
				t.Errorf("a() = %v, want %v", got, tt.want)
			}
		})
	}
}
