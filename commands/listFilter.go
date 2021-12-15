package commands

import (
	"LogAnalyzer/configuration"
	"LogAnalyzer/helper"
	"fmt"
)

// ListFilter lists all available filters
func ListFilter() {
	fmt.Println("Configured filters:")
	maxLength := helper.GetLongestFilterName(configuration.JsonConfig.LogAnalyzer.Filters)

	for i := 0; i < len(configuration.JsonConfig.LogAnalyzer.Filters); i++ {
		indent := helper.GetSpaceSeparator(len(configuration.JsonConfig.LogAnalyzer.Filters[i].Name), maxLength, " ")

		fmt.Println("  " + configuration.JsonConfig.LogAnalyzer.Filters[i].Name + indent +
			"=> Regex: \"" + configuration.JsonConfig.LogAnalyzer.Filters[i].Regex + "\"")
	}
}
