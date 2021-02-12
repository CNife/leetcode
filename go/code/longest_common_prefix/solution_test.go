package longest_common_prefix

import "testing"

//noinspection SpellCheckingInspection
func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty args",
			args: args{},
			want: "",
		},
		{
			name: "has common prefix",
			args: args{[]string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "no common prefix",
			args: args{[]string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "sub slice",
			args: args{[]string{"aa", "a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strings); got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
