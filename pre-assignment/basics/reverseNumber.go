package basics

import (
	"fmt"
	"preassignment/helper"
)

func ReverseTheNumber(value interface{}) (interface{}, error) {
	if val, ok := value.(float32); ok {
		return nil, fmt.Errorf("float values are not allowed %v", val)
	} else if float64val, ok := value.(float64); ok {
		return nil, fmt.Errorf("float values are not allowed %v", float64val)
	}

	if err := helper.IsGivenValueWithInTheRange(value); err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	typeofvalue, err := helper.GetRootDataType(value)
	if err != nil {
		return nil, fmt.Errorf("unsupported data type")
	}

	var int64value int64
	if typeofvalue == "int" {
		if val, ok := value.(uint64); ok {
			var reversenumber int
			for i := int(val); i != 0; i = i / 10 {
				var tempvar int
				tempvar = tempvar % 10
				reversenumber = reversenumber*10 + tempvar
			}
			return interface{}(reversenumber), nil
		}
		int64value = helper.ToInt64(value)
	}

	var reversenumber int
	for i := int(int64value); i != 0; i = i / 10 {
		var tempvar int
		tempvar = i % 10
		reversenumber = reversenumber*10 + tempvar
	}

	return interface{}(reversenumber), nil
}
