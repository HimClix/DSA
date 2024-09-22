package nestedifelses

import "fmt"

func BiggestOfFourNumbersV2(a, b, c, d int) string {
	var biggestNumber int

	if a > b {
		if a > c {
			if a > d {
				biggestNumber = a
			} else {
				biggestNumber = d
			}
		} else {
			if c > d {
				biggestNumber = c
			} else {
				biggestNumber = d
			}
		}
	} else {
		if b > c {
			if b > d {
				biggestNumber = b
			} else {
				biggestNumber = d
			}
		} else {
			if c > d {
				biggestNumber = c
			} else {
				biggestNumber = d
			}
		}
	}

	return fmt.Sprintf("%d is greater than the rest three numbers", biggestNumber)
}
