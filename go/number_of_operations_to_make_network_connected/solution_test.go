package number_of_operations_to_make_network_connected

import "testing"

func TestMakeConnected(t *testing.T) {
	tests := []struct {
		n           int
		connections [][]int
		want        int
	}{
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
	}
	for _, tt := range tests {
		got := MakeConnected(tt.n, tt.connections)
		if got != tt.want {
			t.Errorf("MakeConnected(%v, %v) = %v, want %v",
				tt.n, tt.connections, got, tt.want)
		}
	}
}
