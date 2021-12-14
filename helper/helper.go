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
