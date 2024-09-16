package main

import (
	"fmt"
	"preassignment/basics"
)

func main() {
	//(1a)area of circle
	// radius := 3.144876543456
	// area, err := basics.GetAreaOfCircle(radius)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Printf("%v", area)

	//(1b) swapping of two numbers

	// a := 1
	// b := 2

	// a, b = basics.SwapTwoNumbers(a, b)
	// fmt.Printf("a's value is %d and b's value is %d", a, b)

	//(1c) reverse the given number

	value := interface{}(
		-4294967295890754750,
	)

	val, err := basics.ReverseTheNumber(value)
	if err != nil {
		fmt.Errorf("err %v", err)
	}

	fmt.Printf("reversed number is %v", val)

	//(2b)biggest of two numbers

	// a := int64(777777777777)
	// b := int32(7654567)

	// result, err := controlstructures.FindGreaterNumber(a, b)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Printf("the greater number is %v", result)

}
