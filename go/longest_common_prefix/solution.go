package longest_common_prefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var builder strings.Builder
outer:
	for i := 0; i < len(strs[0]); i++ {
		leftByte := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				break outer
			}
			if rightByte := strs[j][i]; leftByte != rightByte {
				break outer
			}
		}
		builder.WriteByte(leftByte)
	}
	return builder.String()
}
