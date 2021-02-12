package accounts_merge

import (
	"reflect"
	"sort"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		accounts, want [][]string
	}{
		{
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
	}
	for _, tt := range tests {
		got := AccountsMerge(tt.accounts)
		sortStringSlices(got)
		sortStringSlices(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("AccountsMerge(%v) = %v, want %v",
				tt.accounts, got, tt.want)
		}
	}
}

func sortStringSlices(ss [][]string) {
	sort.Slice(ss, func(i, j int) bool {
		if len(ss[i]) != len(ss[j]) {
			return len(ss[i]) < len(ss[j])
		}
		for k := range ss[i] {
			if ss[i][k] != ss[j][k] {
				return ss[i][k] < ss[j][k]
			}
		}
		return false
	})
}
