package ifelse

import (
	"fmt"
	"strings"
)

func VowelIdentifier(a string) string {
	lowercasevalue := strings.ToLower(a)
	if (lowercasevalue == "a") || (lowercasevalue == "e") || (lowercasevalue == "i") || (lowercasevalue == "o") || (lowercasevalue == "u") {
		return fmt.Sprintf("provided character %s is a vowel", a)
	} else {
		return fmt.Sprintf("provided character %s is a ", a)
	}
}
