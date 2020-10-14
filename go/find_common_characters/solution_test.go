package find_common_characters

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		strings []string
		want    []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}
	for _, tt := range tests {
		if got := CommonChars(tt.strings); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CommonChars(%v) = %v, want %v", tt.strings, got, tt.want)
		}
	}
}
