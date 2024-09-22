package ifelse

import "fmt"

func LeapYearIdentifier(year int) string {
	if (year%100 == 0 && year%400 == 0) || year%4 == 0 {
		return fmt.Sprintf("%d is a leap year ", year)
	} else {
		return fmt.Sprintf("%d is a not leap year ", year)
	}
}
