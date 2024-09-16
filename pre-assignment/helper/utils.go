package helper

import (
	"fmt"
	"math"
)

//all ranges values are googled

func IsGivenValueWithInTheRange(val interface{}) error {
	switch v := val.(type) {
	case int8:
		if v < -1<<7 || v > (1<<7)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case int16:
		if v < -1<<15 || v > (1<<15)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case int32:
		if v < -1<<31 || v > (1<<31)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case int64:
		if v < -1<<63 || v > (1<<63)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case int:
		if ^uint(0) == (1<<32 - 1) {
			if v < -1<<31 || v > (1<<31)-1 {
				return fmt.Errorf("given value is out of range")
			}
		} else {
			if v < -1<<63 || v > (1<<63)-1 {
				return fmt.Errorf("given value is out of range")
			}
		}
	case uint8:
		if v > (1<<8)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case uint16:
		if v > (1<<16)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case uint32:
		if v > (1<<32)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case uint64:
		if v > (1<<64)-1 {
			return fmt.Errorf("given value is out of range")
		}
	case uint:
		if ^uint(0) == (1<<32 - 1) {
			if v > (1<<32)-1 {
				return fmt.Errorf("given value is out of range")
			}
		} else {
			if v > (1<<64)-1 {
				return fmt.Errorf("given value is out of range")
			}
		}
	case float32:
		if v < 1.0/(1<<149) || v > (2-2.0/(1<<23))*(1<<127) {
			return fmt.Errorf("given value is out of range")
		}
	case float64:
		//used package here because I was to resolve constant shift overflow issue as value goes too high for max exponent in below commented condition
		// if v < 1.0/(1<<1074) || v > (2 - 2.0/(1<<52))*(1<<1023) {
		// 	return fmt.Errorf("given value is out of range")
		// }
		minFloat64 := math.Pow(2, -1074)
		maxFloat64 := (2 - math.Pow(2, -52)) * math.Pow(2, 1023)
		if v < float64(minFloat64) || v > float64(maxFloat64) {
			return fmt.Errorf("given value is out of range")
		}
	default:
		return fmt.Errorf("invalid data type")
	}
	return nil
}

func GetRootDataType(val interface{}) (string, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64:
		return "int", nil
	case uint, uint8, uint16, uint32, uint64:
		return "int", nil
	case float32, float64:
		return "float", nil
	default:
		return "", fmt.Errorf("invalid data type")
	}
}

func ToInt64(val interface{}) int64 {
	switch v := val.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		if v > math.MaxInt64 {
			panic("uint64 value exceeds int64 range")
		}
		return int64(v)
	default:
		panic("unsupported type for conversion to int64")
	}
}

func ToFloat64(val interface{}) float64 {
	switch v := val.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		panic("unsupported type for conversion to float64")
	}
}
