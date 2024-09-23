package controlstructures

import "fmt"

func TaxCalculator(totalCtc int) string {
	var tax int

	if totalCtc <= 300000 {
		return fmt.Sprintf("Total tax is %d rupees where tax percentage is 0", 0)
	} else if totalCtc <= 700000 {
		tax = (totalCtc - 300000) * 5 / 100
		return fmt.Sprintf("Total tax is %d rupees where tax percentage is 5", tax)
	} else if totalCtc <= 1000000 {
		tax = ((700000 - 300000) * 5 / 100) + ((totalCtc - 700000) * 10 / 100)
		return fmt.Sprintf("Total tax is %d rupees where tax percentage is 10", tax)
	} else if totalCtc <= 1200000 {
		tax = ((700000 - 300000) * 5 / 100) + ((1000000 - 700000) * 10 / 100) + ((totalCtc - 1000000) * 15 / 100)
		return fmt.Sprintf("Total tax is %d rupees where tax percentage is 15", tax)
	} else if totalCtc <= 1500000 {
		tax = ((700000 - 300000) * 5 / 100) + ((1000000 - 700000) * 10 / 100) + ((1200000 - 1000000) * 15 / 100) + ((totalCtc - 1200000) * 20 / 100)
		return fmt.Sprintf("Total tax is %d rupees where tax percentage is 20", tax)
	} else {
		tax = ((700000 - 300000) * 5 / 100) + ((1000000 - 700000) * 10 / 100) + ((1200000 - 1000000) * 15 / 100) + ((1500000 - 1200000) * 20 / 100) + ((totalCtc - 1500000) * 30 / 100)
		return fmt.Sprintf("Total tax is %d rupees where tax percentage is 30", tax)
	}
}
