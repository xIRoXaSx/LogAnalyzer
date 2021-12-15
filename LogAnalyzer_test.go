package main

import (
	"LogAnalyzer/configuration"
	"LogAnalyzer/logger"
	"os"
	"path/filepath"
	"testing"
)

var exampleLogPath = filepath.Join(configuration.ConfigPath, "example.log")

func TestCreateConfig(t *testing.T) {
	if _, err := os.Stat(configuration.ConfigPath); err == nil {
		if err := os.Remove(configuration.ConfigFullPath); err != nil {
			logger.Critical("Cannot remove config folder!")
			panic("Cannot remove config folder")
		}
	}

	createExampleLog()
	configuration.CreateConfigIfNotExists()
	configuration.ReadJson()
	CheckArgs()
}

func TestListFilter(t *testing.T) {
	os.Args = []string{"PathOfBin", "listfilter"}
	configuration.ReadJson()
	CheckArgs()
}

func TestInspectInfoSuccess(t *testing.T) {
	os.Args = []string{"PathOfBin", "ins", "Info", exampleLogPath}
	main()
}

func TestTooLessArgs(t *testing.T) {
	os.Args = []string{"PathOfBin"}
	main()
}

func TestRemoveLog(t *testing.T) {
	deleteExampleLog()
}

func createExampleLog() {
	logContent := []byte(`
[00:00:00] [Info] This is a test information log entry!
[00:00:00] Nothing special, nothing fancy...
[00:00:00] [Error] This is a test error log entry!
`)
	if err := os.WriteFile(exampleLogPath, logContent, 0700); err != nil {
		logger.Error("Cannot write example log (" + exampleLogPath + ")!")
	}
}

func deleteExampleLog() {
	if err := os.Remove(exampleLogPath); err != nil {
		logger.Error("Cannot delete example log (" + exampleLogPath + ")!")
	}
}
