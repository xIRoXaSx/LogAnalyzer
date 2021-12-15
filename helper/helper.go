package helper

import (
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"os"
	"strings"
)

func Contains(slice []structs.Filter, desired string) structs.Filter {
	returnValue := structs.Filter{}

	for i := 0; i < len(slice); i++ {
		if strings.EqualFold(desired, slice[i].Name) {
			returnValue = slice[i]
			break
		}
	}

	return returnValue
}

// GetFileContent gets the content of filePath from the passed string
func GetFileContent(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error(err.Error())
		return []byte{}
	}

	return content
}

// GetLongestFilterName gets the longest filter name of structs.Filter in a slice
func GetLongestFilterName(textSlice []structs.Filter) int {
	maxLength := 0
	for i := 0; i < len(textSlice); i++ {
		if maxLength < len(textSlice[i].Name) {
			maxLength = len(textSlice[i].Name)
		}
	}

	return maxLength
}

// GetLongestStringCommand gets the longest string in a command slice
func GetLongestStringCommand(command []structs.Command) (int, int, int) {
	maxLengthName := 0
	maxLengthUsage := 0
	maxLengthDescription := 0

	for i := 0; i < len(command); i++ {
		if len(command[i].Name) > maxLengthName {
			maxLengthName = len(command[i].Name)
		}

		if len(command[i].Usage) > maxLengthUsage {
			maxLengthUsage = len(command[i].Usage)
		}

		if len(command[i].Description) > maxLengthDescription {
			maxLengthDescription = len(command[i].Description)
		}
	}

	return maxLengthName, maxLengthUsage, maxLengthDescription
}

// GetSpaceSeparator gets the spaces needed for maxLength - length + 1
func GetSpaceSeparator(length int, maxLength int, separator string) string {
	var calcLen int
	if maxLength > length {
		calcLen = maxLength - length
	} else {
		calcLen = length - maxLength
	}

	return strings.Repeat(separator, calcLen+5)
}
