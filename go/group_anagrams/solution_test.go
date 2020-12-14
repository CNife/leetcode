package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strings []string
		want    [][]string
	}{
		{
			strings: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		got := GroupAnagrams(tt.strings)
		got, want := sortStringSlices(got), sortStringSlices(tt.want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strings, got, want)
		}
	}
}

func sortStringSlices(src [][]string) [][]string {
	for _, slice := range src {
		sort.Strings(slice)
	}
	sort.Slice(src, func(i, j int) bool {
		m, n := 0, 0
		for m < len(src[i]) && n < len(src[j]) {
			lhs, rhs := src[i][m], src[j][n]
			if lhs != rhs {
				return lhs < rhs
			}
			m++
			n++
		}
		return m == len(src[i]) && n < len(src[j])
	})
	return src
}
