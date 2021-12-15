package structs

// JsonConf is the root object inside the json config
type JsonConf struct {
	LogAnalyzer LogAnalyzer
}

// LogAnalyzer is the main object
type LogAnalyzer struct {
	EnableDebug bool
	Filters     []Filter
}

// Filter is the array / slice of all filter objects
type Filter struct {
	Name  string
	Regex string
}

// Command is a struct for each available command with its Usage and Description
type Command struct {
	Name        string
	Usage       string
	Description string
}
