package controlstructures

import "fmt"

func BiggestOfFourNumbers(a, b, c, d int) string {

	if a > b && a > c && a > d {
		return fmt.Sprintf("%d is greater than %d,%d,%d", a, b, c, d)
	} else if b > a && b > c && b > d {
		return fmt.Sprintf("%d is greater than %d,%d,%d", b, a, c, d)
	} else if c > a && c > b && c > d {
		return fmt.Sprintf("%d is greater than %d,%d,%d", c, a, b, d)
	} else if d > a && d > b && d > c {
		return fmt.Sprintf("%d is greater than %d,%d,%d", d, a, b, c)
	}

	return "something went wrong"
}
