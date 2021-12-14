package commands

import (
	"LogAnalyzer/helper"
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"fmt"
	"regexp"
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
	matched, err := regexp.MatchString(filter.Regex, text)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if !matched {
		logger.Info("No match found!")
		return
	}

	matchedStrings := getRegexMatches(text, filter.Regex)
	if len(matchedStrings) < 1 {
		return
	}

	for i := 0; i < len(matchedStrings); i++ {
		fmt.Println(matchedStrings[i])
	}

	fmt.Println()
	logger.Info("Finished after " + (time.Since(start).Round(time.Millisecond).String()))
}

// getRegexMatches gets the regex matches from the given parameters
func getRegexMatches(text string, regexString string) []string {
	regex := regexp.MustCompile(regexString)
	return regex.FindAllString(text, -1)
}
