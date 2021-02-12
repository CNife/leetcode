package queue_reconstruction_by_height

import (
	"reflect"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	tests := []struct {
		people, want [][]int
	}{
		{
			people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			want:   [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
	}
	for _, tt := range tests {
		if got := ReconstructQueue(tt.people); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ReconstructQueue(%v) = %v, want %v", tt.people, got, tt.want)
		}
	}
}
