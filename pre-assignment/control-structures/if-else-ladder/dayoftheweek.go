package controlstructures

import "fmt"

func IdentifyDayOfTheWeek(day, month, year int) string {
	//solving using zeller's formula
	//h=[d+[13(m+1)/5]+k+(k/4)+(J/4)-(2J)]%7
	// d,m,y = date,month,year
	//k=year of the century(y%100)
	//y= century(y/100)
	//according zeller 0,1 are sat and sun
	// 2 to 6 are mon to fri

	dayoftheweek := (day + ((13 * (month + 1)) / 5) + (year % 100) + ((year % 100) / 4) + ((year / 100) / 4) - (2 * (year / 100))) % 7
	if dayoftheweek == 0 {
		return fmt.Sprintf("saturday is day of %d/%d/%d", day, month, year)
	} else if dayoftheweek == 1 {
		return fmt.Sprintf("sunday is day of %d/%d/%d", day, month, year)
	} else if dayoftheweek == 2 {
		return fmt.Sprintf("monday is day of %d/%d/%d", day, month, year)
	} else if dayoftheweek == 3 {
		return fmt.Sprintf("tuesday is day of %d/%d/%d", day, month, year)
	} else if dayoftheweek == 4 {
		return fmt.Sprintf("wednesday is day of %d/%d/%d", day, month, year)
	} else if dayoftheweek == 5 {
		return fmt.Sprintf("thursday is day of %d/%d/%d", day, month, year)
	} else {
		return fmt.Sprintf("friday is day of %d/%d/%d", day, month, year)
	}
}
