package rectangle_area

func ComputeArea(a, b, c, d, e, f, g, h int) int {
	area1, area2 := area(a, b, c, d), area(e, f, g, h)
	overlapWidth, overlapHeight := overlapLength(a, c, e, g), overlapLength(b, d, f, h)
	return area1 + area2 - overlapWidth*overlapHeight
}

func area(a, b, c, d int) int {
	return (c - a) * (d - b)
}

func overlapLength(a, b, c, d int) int {
	if a > c {
		a, b, c, d = c, d, a, b
	}
	if b <= c {
		return 0
	}
	if b < d {
		return b - c
	}
	return d - c
}
