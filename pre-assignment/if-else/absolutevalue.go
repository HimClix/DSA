package ifelse

func AbsoluteValueOfTheNumber(a int) int {
	if a < 0 {
		a = -a
		return a
	} else {
		return a
	}
}
