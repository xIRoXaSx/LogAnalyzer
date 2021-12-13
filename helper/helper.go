package helper

import (
	"LogAnalyzer/configuration"
	"strings"
)

func Contains(slice []configuration.Filter, desired string) configuration.Filter {
	returnValue := configuration.Filter{}

	for i := 0; i < len(slice); i++ {
		if strings.EqualFold(desired, slice[i].Name) {
			returnValue = slice[i]
			break
		}
	}

	return returnValue
}
