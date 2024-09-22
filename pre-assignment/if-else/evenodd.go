package ifelse

import "fmt"

func FindEvenOrOdd(a int) string {
	if a%2 == 0 {
		return fmt.Sprintf("provided value %d is even", a)
	} else {
		return fmt.Sprintf("provided value %d is odd", a)
	}
}
