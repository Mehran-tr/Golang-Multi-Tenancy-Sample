package utils

import "log"

// LogInfo logs informational messages
func LogInfo(message string) {
	log.Printf("[INFO]: %s\n", message)
}

// LogError logs error messages
func LogError(err error) {
	log.Printf("[ERROR]: %v\n", err)
}
