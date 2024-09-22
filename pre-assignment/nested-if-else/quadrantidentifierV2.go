package nestedifelses

import "fmt"

func QuadrantPointIdentifierV2(x, y int) string {
	if x >= 0 {
		if y >= 0 {
			return fmt.Sprintf("provided point %d,%d exists in 1st quadrant", x, y)
		} else {
			return fmt.Sprintf("provided point %d,%d exists in 4th quadrant", x, y)
		}
	} else {
		if x < 0 {
			if y >= 0 {
				return fmt.Sprintf("provided point %d,%d exists in 2nd quadrant", x, y)
			}
		} else {
			return fmt.Sprintf("provided point %d,%d exists in 3rd quadrant", x, y)
		}
	}

	return "something went wrong"
}
