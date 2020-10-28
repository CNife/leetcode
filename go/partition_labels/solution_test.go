package partition_labels

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		src  string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}
	for _, tt := range tests {
		if got := PartitionLabels(tt.src); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PartitionLabels(%v) = %v, want %v", tt.src, got, tt.want)
		}
	}
}
