package controlstructures

import "fmt"

func GradeBasedOnAverageMarks(totalsubjects, totalmarks int) string {
	percentage := totalmarks / totalsubjects

	if percentage >= 90 {
		return fmt.Sprintf("the grade is %s", "A+")
	} else if percentage >= 75 && percentage < 90 {
		return fmt.Sprintf("the grade is %s", "A")
	} else if percentage >= 60 && percentage < 75 {
		return fmt.Sprintf("the grade is %s", "B")
	} else if percentage >= 45 && percentage < 60 {
		return fmt.Sprintf("the grade is %s", "C")
	} else {
		return fmt.Sprintf("the grade is %s", "D")
	}
}
