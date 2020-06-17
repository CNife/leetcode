package best_sightseeing_pair

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[]int{8, 1, 5, 2, 6}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxScoreSightseeingPair(tt.args.array); got != tt.want {
				t.Errorf("MaxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
