package basics

import (
	"fmt"
	"preassignment/helper"
)

func GetAreaOfCircle(radius interface{}) (interface{}, error) {
	valueofpie := 3.141592653589793

	typeofradius, err := helper.GetRootDataType(radius)
	if err != nil {
		return nil, fmt.Errorf("invalid data type for 'a': %s", err.Error())
	}
	var area interface{}
	if typeofradius == "int" {
		intradius := helper.ToInt64(radius)
		area = valueofpie * float64((intradius * intradius))
	} else {
		floatradius := helper.ToFloat64(radius)
		area = float64(valueofpie) * (floatradius * floatradius)
	}

	return area, nil
}
