package lemonade_change

func LemonadeChange(bills []int) bool {
	count5, count10, count20 := 0, 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			count5++
		case 10:
			if count5 > 0 {
				count5, count10 = count5-1, count10+1
			} else {
				return false
			}
		case 20:
			if count10 > 0 && count5 > 0 {
				count5, count10, count20 = count5-1, count10-1, count20+1
			} else if count5 > 2 {
				count5, count20 = count5-3, count20+1
			} else {
				return false
			}
		default:
			panic("invalid bill")
		}
	}
	return true
}
