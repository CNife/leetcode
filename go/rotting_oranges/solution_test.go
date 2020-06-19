package rotting_oranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}},
			want: 4,
		},
		{
			name: "2",
			args: args{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}},
			want: -1,
		},
		{
			name: "3",
			args: args{[][]int{{0, 2}}},
			want: 0,
		},
		{
			name: "4",
			args: args{[][]int{{0}}},
			want: 0,
		},
		{
			name: "5",
			args: args{[][]int{{1}}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("OrangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
