package main

const (
	ConfigDir  = "LogAnalyzer"
	ConfigFile = "config"
)

type Config struct {
	Debug      bool
	PrintStats bool
	Filters    []Filter
}

var conf *Config

func defaultConfig() *Config {
	return &Config{
		Debug:      false,
		PrintStats: true,
		Filters: []Filter{
			{
				Name:  "JsonMin",
				Regex: "(\\s+[^{}\"'\\[\\]\\\\\\w])|(\\B\\s)",
				Options: &FilterOpts{
					Replacement:      "",
					RemoveEmptyLines: true,
				},
			},
			{
				Name:  "JavaStackTrace",
				Regex: "(?m)^.*?Exception.*(?:[\\r|\\n]+^\\s*at .*)+",
			},
			{
				Name:  "StackTrace",
				Regex: "(?m)((.*(\\n|\\r|\\r\\n)){1})^.*?Exception.*(?:[\\n|\\r|\\r\\n]+^\\s*at .*)+",
				Options: &FilterOpts{
					Replacement:      "Nope, not a single error to report ;)",
					RemoveEmptyLines: true,
				},
			},
		},
	}
}
