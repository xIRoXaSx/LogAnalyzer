package commands

import (
	"LogAnalyzer/helper"
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"fmt"
	"time"
)

// Inspect checks if the given filePath has any line matching the provided regex
func Inspect(filePath string, filter structs.Filter) {
	start := time.Now()

	if filter == (structs.Filter{}) {
		logger.Error("Passed filter is empty!")
		return
	}

	returnValue := helper.GetFileContent(filePath)
	if len(returnValue) < 1 {
		return
	}

	text := string(returnValue)
	matchedStrings := helper.GetAllRegexpMatches(text, filter.Regex)

	if len(matchedStrings) < 1 {
		return
	}

	for i := 0; i < len(matchedStrings); i++ {
		fmt.Println(matchedStrings[i])
	}

	fmt.Println()
	logger.Info("Finished after " + helper.CalculateExecutionDuration(start))
}
