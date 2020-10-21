package long_pressed_name

import "testing"

func TestIsLongPressedName(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		name, typed string
		want        bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
		{"vtkgn", "vttkgnn", true},
	}
	for _, tt := range tests {
		if got := IsLongPressedName(tt.name, tt.typed); got != tt.want {
			t.Errorf("IsLongPressedName(%v, %v) = %v, want %v", tt.name, tt.typed, got, tt.want)
		}
	}
}
