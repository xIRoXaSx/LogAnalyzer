package configuration

import (
	"LogAnalyzer/logger"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type empty struct{}

var configFileName = "config.json"
var configFolderName = packageName[:strings.IndexByte(packageName, '/')]
var configBasePath, _ = os.UserConfigDir()
var configPath = filepath.Join(configBasePath, configFolderName)
var packageName = reflect.TypeOf(empty{}).PkgPath()
var configFullPath = filepath.Join(configBasePath, configFolderName, configFileName)

// CreateConfigIfNotExists creates / copies the default configuration if it does not exist locally
func CreateConfigIfNotExists() {
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configPath, 0700)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		logger.Info("Created config directory \"" + configPath + "\"")
	}

	if _, err := os.Stat(configFullPath); err == nil {
		return
	}

	copyFile()
	return
}

// CopyFile copies the content of the default config.json to the local one
func copyFile() {
	_, filename, _, gotInfo := runtime.Caller(0)
	if !gotInfo {
		panic("No caller information")
	}

	// Open the config configFile
	readBytes, err := ioutil.ReadFile(filepath.Join(path.Dir(filename), configFileName))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if len(readBytes) < 1 {
		return
	}

	// Open the config configFile
	if err := ioutil.WriteFile(configFullPath, readBytes, 0700); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Created config file \"" + configFullPath + "\"")
	}
}
