package controlstructures

func ActionBasedOperations(a, b int, action string) int {
	var c int
	if action == "add" {
		c = a + b
	} else if action == "substract" {
		c = a - b
	} else if action == "divison" {
		c = a / b
	} else {
		c = a * b
	}
	return c
}
