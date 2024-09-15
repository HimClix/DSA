package controlstructures

import (
	"fmt"
	"preassignment/helper"
)

func FindGreaterNumber(a, b interface{}) (interface{}, error) {
	if err := helper.IsGivenValueWithInTheRange(a); err != nil {
		return nil, fmt.Errorf("given 1st parameter is out of range: %s", err.Error())
	}

	if err := helper.IsGivenValueWithInTheRange(b); err != nil {
		return nil, fmt.Errorf("given 2nd parameter is out of range: %s", err.Error())
	}

	typeofa, err := helper.GetRootDataType(a)
	if err != nil {
		return nil, fmt.Errorf("invalid data type for 'a': %s", err.Error())
	}

	typeofb, err := helper.GetRootDataType(b)
	if err != nil {
		return nil, fmt.Errorf("invalid data type for 'b': %s", err.Error())
	}

	if typeofa != typeofb {
		return nil, fmt.Errorf("cannot compare different data types: '%s' vs '%s'", typeofa, typeofb)
	}

	if typeofa == "int" && typeofb == "int" {

		va, vb := helper.ToInt64(a), helper.ToInt64(b)

		if va > vb {
			return a, nil
		} else {
			return b, nil
		}

	}

	if typeofa == "float" && typeofb == "float" {

		va, vb := helper.ToFloat64(a), helper.ToFloat64(a)

		if va > vb {
			return a, nil
		} else {
			return b, nil
		}

	}

	return nil, fmt.Errorf("unsupported data type for comparison")
}
