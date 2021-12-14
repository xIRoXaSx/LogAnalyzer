package commands

import (
	"LogAnalyzer/configuration"
	"LogAnalyzer/helper"
	"fmt"
	"strings"
)

func ListFilter() {
	fmt.Println("Configured filters:")
	maxLength := helper.GetLongestFilterName(configuration.JsonConfig.LogAnalyzer.Filters)

	for i := 0; i < len(configuration.JsonConfig.LogAnalyzer.Filters); i++ {
		charsToIndent := maxLength - len(configuration.JsonConfig.LogAnalyzer.Filters[i].Name) + 1
		fmt.Println("  " + configuration.JsonConfig.LogAnalyzer.Filters[i].Name + strings.Repeat(" ", charsToIndent) +
			"=> Regex: \"" + configuration.JsonConfig.LogAnalyzer.Filters[i].Regex + "\"")
	}
}
