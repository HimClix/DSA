package main

import (
	"fmt"
	"preassignment/basics"
)

func main() {
	//(1a)area of circle
	radius := 3.144876543456
	area, err := basics.GetAreaOfCircle(radius)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("%v", area)

	//(2b)biggest of two numbers

	// a := int64(777777777777)
	// b := int32(7654567)

	// result, err := controlstructures.FindGreaterNumber(a, b)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Printf("the greater number is %v", result)

}
