package dota2_senate

const (
	radiant = "Radiant"
	dire    = "Dire"
)

func PredicatePartyVictory(senate string) string {
	rNumber, dNumber := 0, 0
	curBanR, curBanD := 0, 0
	totalBanR, totalBanD := 0, 0
	flag := true
	bytes := []byte(senate)
	for {
		for i := 0; i < len(bytes); i++ {
			switch bytes[i] {
			case 'R':
				if flag {
					rNumber++
				}
				if curBanR == 0 {
					curBanD++
					totalBanD++
					if totalBanD == dNumber && !flag {
						return radiant
					}
				} else {
					curBanR--
					bytes[i] = 'r'
				}
			case 'D':
				if flag {
					dNumber++
				}
				if curBanD == 0 {
					curBanR++
					totalBanR++
					if totalBanR == rNumber && !flag {
						return dire
					}
				} else {
					curBanD--
					bytes[i] = 'd'
				}
			}
		}
		flag = false
		if totalBanD >= dNumber {
			return radiant
		}
		if totalBanR >= rNumber {
			return dire
		}
	}
}
