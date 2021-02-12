package surface_area_of_3d_shapes

import "testing"

func Test_surfaceArea(t *testing.T) {
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
			args: args{[][]int{{2}}},
			want: 10,
		},
		{
			name: "2",
			args: args{[][]int{{1, 2}, {3, 4}}},
			want: 34,
		},
		{
			name: "3",
			args: args{[][]int{{1, 0}, {0, 2}}},
			want: 16,
		},
		{
			name: "4",
			args: args{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			want: 32,
		},
		{
			name: "5",
			args: args{[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
			want: 46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SurfaceArea(tt.args.grid); got != tt.want {
				t.Errorf("SurfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
