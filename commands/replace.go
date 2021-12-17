package commands

import (
	"LogAnalyzer/helper"
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"fmt"
	"regexp"
	"time"
)

// Replace uses the passed filter to replace the matched strings in the file
func Replace(filePath string, filter structs.Filter, replacement string) {
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

	regex := regexp.MustCompile(filter.Regex)
	replaced := regex.ReplaceAllString(text, replacement)

	if filter.RemoveEmptyLines {
		regexNewLine := regexp.MustCompile("[\n]+")
		regexReturnNewLine := regexp.MustCompile("[\r\n]+")
		replaced = regexNewLine.ReplaceAllString(replaced, "\n")
		replaced = regexReturnNewLine.ReplaceAllString(replaced, "\r\n")
	}

	fmt.Println(replaced)

	if filter.DontPrintStats {
		return
	}

	fmt.Println()
	logger.Info("Finished after " + helper.CalculateExecutionDuration(start))
}
