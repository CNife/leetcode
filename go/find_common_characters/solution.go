package find_common_characters

const tableSize = 26

type table = [tableSize]uint

func CommonChars(strings []string) []string {
	if len(strings) == 0 {
		return nil
	}

	minTable := buildTable(strings[0])
	for i := 1; i < len(strings); i++ {
		tmpTable := buildTable(strings[i])
		for j := 0; j < tableSize; j++ {
			if minTable[j] > tmpTable[j] {
				minTable[j] = tmpTable[j]
			}
		}
	}

	var result []string
	for i := 0; i < tableSize; i++ {
		for j := uint(0); j < minTable[i]; j++ {
			result = append(result, string(byte(i)+'a'))
		}
	}
	return result
}

func buildTable(s string) table {
	var result table
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		result[index]++
	}
	return result
}
