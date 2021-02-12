package dota2_senate

import "testing"

func TestPredicatePartyVictory(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		senate, want string
	}{
		{"RDD", dire},
		{"RD", radiant},
		{"DRDRR", dire},
	}
	for _, tt := range tests {
		if got := PredicatePartyVictory(tt.senate); got != tt.want {
			t.Errorf("PredicatePartyVictory(%v) = %v, want %v",
				tt.senate, got, tt.want)
		}
	}
}
