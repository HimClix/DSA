package controlstructures

import "fmt"

func LeapYearIdentifierV2(year int) string {
	if year%100 == 0 && year%400 == 0 {
		return fmt.Sprintf("%d is a leap year ", year)
	} else if year%100 == 0 {
		return fmt.Sprintf("%d is not a leap year ", year)
	} else if year%4 == 0 {
		return fmt.Sprintf("%d is a leap year ", year)
	} else {
		return fmt.Sprintf("%d is not a leap year ", year)
	}
}
