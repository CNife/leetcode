package isomorphic_strings

func IsIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) < 1 {
		return true
	}

	var charTab, revCharTab [128]byte
	for i := range s {
		cs, ct := s[i], t[i]
		if charTab[cs] == 0 && revCharTab[ct] == 0 {
			charTab[cs] = ct
			revCharTab[ct] = cs
		} else if charTab[cs] != ct || revCharTab[ct] != cs {
			return false
		}
	}
	return true
}
