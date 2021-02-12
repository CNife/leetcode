package repeated_dna_sequences

import (
	"reflect"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		src  string
		want []string
	}{
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
	}
	for _, tt := range tests {
		if got := FindRepeatedDnaSequences(tt.src); !reflect.DeepEqual(mapper(got), mapper(tt.want)) {
			t.Errorf("TestFindRepeatedDnaSequences(%v) = %v, want %v", tt.src, got, tt.want)
		}

	}
}

func mapper(ss []string) map[string]int {
	m := make(map[string]int, len(ss))
	for _, s := range ss {
		if count, exists := m[s]; exists {
			m[s] = count + 1
		} else {
			m[s] = 1
		}
	}
	return m
}
