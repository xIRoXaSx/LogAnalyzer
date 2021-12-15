package configuration

import (
	"LogAnalyzer/logger"
	"LogAnalyzer/structs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var JsonConfig structs.JsonConf
var configFileName = "config.json"
var configFolderName = packageName[:strings.IndexByte(packageName, '/')]
var configBasePath, _ = os.UserConfigDir()
var ConfigPath = filepath.Join(configBasePath, configFolderName)
var packageName = reflect.TypeOf(structs.JsonConf{}).PkgPath()
var ConfigFullPath = filepath.Join(configBasePath, configFolderName, configFileName)

// CreateConfigIfNotExists creates / copies the default configuration if it does not exist locally
func CreateConfigIfNotExists() {
	if _, err := os.Stat(ConfigPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(ConfigPath, 0700)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		logger.Info("Created config directory \"" + ConfigPath + "\"")
	}

	if _, err := os.Stat(ConfigFullPath); err == nil {
		return
	}

	writeConfig()
	return
}

// writeConfig writes the config file to the config path
func writeConfig() {
	jsonConfig := structs.JsonConf{
		LogAnalyzer: structs.LogAnalyzer{
			EnableDebug: false,
			Filters: []structs.Filter{
				{Name: "Info", Regex: "(?m)^.*\\[.*INFO\\].*"},
				{Name: "Error", Regex: "(?m)^.*\\[.*ERROR\\].*"},
				{Name: "StackTrace", Regex: "(?m)((.*(\\n|\\r|\\r\\n)){1})^.*?Exception.*(?:[\\n|\\r|\\r\\n]+^\\s*at .*)+"},
				{Name: "JavaStackTrace", Regex: "(?m)^.*?Exception.*(?:[\\r|\\n]+^\\s*at .*)+"},
			},
		},
	}

	content, err := json.MarshalIndent(jsonConfig, "", "\t")
	if err != nil {
		logger.Error(err.Error())
	}

	if err := ioutil.WriteFile(ConfigFullPath, content, 0700); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Created config file \"" + ConfigFullPath + "\"")
	}
}

// ReadJson unmarshalls the json config file
func ReadJson() {
	jsonFile, err := os.Open(ConfigFullPath)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			logger.Error(err.Error())
			panic("Cannot close file")
		}
	}(jsonFile)

	err = json.Unmarshal(bytes, &JsonConfig)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if !JsonConfig.LogAnalyzer.EnableDebug {
		return
	}

	for i := 0; i < len(JsonConfig.LogAnalyzer.Filters); i++ {
		logger.Debug("Filter \"" + JsonConfig.LogAnalyzer.Filters[i].Name +
			"\" loaded with Regex \"" + JsonConfig.LogAnalyzer.Filters[i].Regex + "\"",
		)
	}
}
