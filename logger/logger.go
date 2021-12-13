package logger

import (
	"fmt"
	"time"
)

var reset = "\033[0m"
var red = "\033[31m"
var yellow = "\033[33m"
var cyan = "\033[36m"
var white = "\033[37m"

// Debug logs debug messages to the console
func Debug(text string) {
	fmt.Println(cyan+"DBG:", GetCurrentTime(), "\t"+text+reset)
}

// Info logs informational messages to the console
func Info(text string) {
	fmt.Println(white+"INF:", GetCurrentTime(), "\t"+text+reset)
}

// Warn logs warn messages to the console
func Warn(text string) {
	fmt.Println(yellow+"WRN:", GetCurrentTime(), "\t"+text+reset)
}

// Error logs warn messages to the console
func Error(text string) {
	fmt.Println(red+"ERR:", GetCurrentTime(), "\t"+text+reset)
}

// Critical logs warn messages to the console
func Critical(text string) {
	fmt.Println(red+"CRT:", GetCurrentTime(), "\t"+text+reset)
}

// GetCurrentTime gets the current time formatted to "YYYY-MM-DD HH:MM:SS"
func GetCurrentTime() string {
	now := time.Now()
	return fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
}
